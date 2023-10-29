package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/meetmorrowsolonmars/zhopij/answer/internal/app/v1/debug"
	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pkg/consumers"
	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pkg/producers"
	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pkg/repositories"
	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pkg/services"
)

func main() {
	const (
		ServerAddressGRPC  = ":82"
		ServerAddressDebug = ":84"
	)

	// logger setup
	logger := zap.Must(zap.NewProductionConfig().Build()).Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	// metrics setup
	metrics := grpcprom.NewServerMetrics(grpcprom.WithServerHandlingTimeHistogram())
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	// graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// kafka setup
	// TODO: move to yaml config
	const (
		broker               = "kafka:29092"
		consumerGroupID      = "answer"
		answerEventTopicName = "answer_events_test"
	)

	saramaConfig := sarama.NewConfig()
	saramaConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategyRoundRobin(),
		sarama.NewBalanceStrategyRange(),
	}
	saramaConfig.Consumer.Offsets.Initial = sarama.OffsetNewest
	saramaConfig.Consumer.Return.Errors = false
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.Return.Errors = true

	echoConsumerGroup, err := sarama.NewConsumerGroup([]string{broker}, consumerGroupID+"_echo", saramaConfig)
	if err != nil {
		logger.Fatalf("can't create consumer group: %s", err)
	}

	consumerGroup, err := sarama.NewConsumerGroup([]string{broker}, consumerGroupID, saramaConfig)
	if err != nil {
		logger.Fatalf("can't create consumer group: %s", err)
	}

	syncProducer, err := sarama.NewSyncProducer([]string{broker}, saramaConfig)
	if err != nil {
		logger.Fatalf("can't create sync producer: %s", err)
	}

	// database connection setup
	// TODO: move to yaml config
	user := os.Getenv("POSTGRES_ANSWER_USER")
	password := os.Getenv("POSTGRES_ANSWER_USER_PASSWORD")
	name := os.Getenv("POSTGRES_ANSWER_DB")
	host := os.Getenv("POSTGRES_CONTAINER_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, name, host, port)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Fatalf("can't parse database config: %s", err)
	}

	config.MaxConns = 40
	config.MinConns = 10

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		logger.Fatalf("can't create database connection: %s", err)
	}

	// domain services setup
	marshalOptions := protojson.MarshalOptions{
		UseProtoNames:  true,
		UseEnumNumbers: false,
	}
	unmarshalOptions := protojson.UnmarshalOptions{
		AllowPartial:   false,
		DiscardUnknown: true,
	}

	answerRepository := repositories.NewAnswerRepository(3*time.Second, db)
	answerService := services.NewAnswerService(answerRepository)

	echoConsumer := consumers.NewEchoConsumer(logger)
	answerEventsConsumer := consumers.NewAnswerEventsConsumer(answerService, logger, unmarshalOptions)
	eventsProducer := producers.NewAnswerEventsProducer(answerEventTopicName, syncProducer, marshalOptions)

	// grpc services setup
	debugGrpcServer := debug.NewDebugServiceImplementation(eventsProducer)

	// grpc server setup
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(metrics.UnaryServerInterceptor()))

	debug.RegisterGRPCServer(server, debugGrpcServer)

	metrics.InitializeMetrics(server)
	reflection.Register(server)

	go func() {
		<-ctx.Done()
		server.Stop()
	}()

	// kafka start consumer group
	go func() {
		for ctx.Err() == nil {
			err := echoConsumerGroup.Consume(ctx, []string{answerEventTopicName}, echoConsumer)
			if err != nil {
				logger.Errorf("consumer group consume err: %s", err)
			}
		}
	}()

	go func() {
		for ctx.Err() == nil {
			err := consumerGroup.Consume(ctx, []string{answerEventTopicName}, answerEventsConsumer)
			if err != nil {
				logger.Errorf("consumer group consume err: %s", err)
			}
		}
	}()

	// debug server startup
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/debug", promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}))
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

		if err := http.ListenAndServe(ServerAddressDebug, mux); err != nil {
			logger.Fatalf("can't start app: %s", err)
		}
	}()

	// grpc server startup
	logger.Infof("Server startup on port: %s", ServerAddressGRPC)

	lis, err := net.Listen("tcp", ServerAddressGRPC)
	if err != nil {
		logger.Fatalf("can't start app: %s", err)
	}

	if err = server.Serve(lis); err != nil {
		logger.Fatalf("can't start app: %s", err)
	}
}
