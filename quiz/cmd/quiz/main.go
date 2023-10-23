package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/meetmorrowsolonmars/zhopij/quiz/internal/app/v1/quiz"
	"github.com/meetmorrowsolonmars/zhopij/quiz/internal/pkg/repositories"
	"github.com/meetmorrowsolonmars/zhopij/quiz/internal/pkg/services"
)

func main() {
	const (
		ServerAddressGRPC  = ":82"
		ServerAddressDebug = ":84"
	)

	ctx := context.Background()

	// logger setup
	logger := zap.Must(zap.NewProductionConfig().Build()).Sugar()
	defer func() {
		_ = logger.Sync()
	}()

	// metrics setup
	metrics := grpcprom.NewServerMetrics(grpcprom.WithServerHandlingTimeHistogram())
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	// database connection setup
	// TODO: move to yaml config
	user := os.Getenv("POSTGRES_QUIZ_USER")
	password := os.Getenv("POSTGRES_QUIZ_USER_PASSWORD")
	name := os.Getenv("POSTGRES_QUIZ_DB")
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
	quizRepository := repositories.NewQuizRepository(3*time.Second, db)
	quizService := services.NewQuizService(quizRepository)

	// grpc services setup
	quizGrpcServer := quiz.NewQuizServiceImplementation(quizService, logger)

	// grpc server setup
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(metrics.UnaryServerInterceptor()))

	quiz.RegisterGRPCServer(server, quizGrpcServer)

	metrics.InitializeMetrics(server)
	reflection.Register(server)

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
