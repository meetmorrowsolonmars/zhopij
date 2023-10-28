package main

import (
	"context"
	"net"
	"net/http"
	"net/http/pprof"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/meetmorrowsolonmars/zhopij/answer/internal/app/v1/debug"
	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pkg/consumers"
	"github.com/meetmorrowsolonmars/zhopij/answer/internal/pkg/producers"
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
	// TODO: move to config
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
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.Return.Errors = true

	consumerGroup, err := sarama.NewConsumerGroup([]string{broker}, consumerGroupID, saramaConfig)
	if err != nil {
		logger.Fatalf("can't create consumer group: %s", err)
	}

	syncProducer, err := sarama.NewSyncProducer([]string{broker}, saramaConfig)
	if err != nil {
		logger.Fatalf("can't create sync producer: %s", err)
	}

	// domain services setup
	echoConsumer := consumers.NewEchoConsumer(logger)
	eventsProducer := producers.NewAnswerEventsProducer(answerEventTopicName, syncProducer)

	// grpc services setup
	debugGrpcServer := debug.NewDebugServiceImplementation(eventsProducer, protojson.MarshalOptions{
		UseProtoNames:  true,
		UseEnumNumbers: false,
	})

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
		err := consumerGroup.Consume(ctx, []string{answerEventTopicName}, echoConsumer)
		if err != nil {
			logger.Fatalf("can't start consumer group: %s", err)
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
