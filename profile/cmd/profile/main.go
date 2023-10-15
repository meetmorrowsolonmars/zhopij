package main

import (
	profile2 "github.com/meetmorrowsolonmars/zhopij/profile/internal/app/v1/profile"
	"github.com/meetmorrowsolonmars/zhopij/profile/internal/pb/api/v1/profile"
	"net"

	grpcprom "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// grpc server setup
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		metrics.UnaryServerInterceptor(),
	))

	profile.RegisterProfileServiceServer(server, profile2.NewProfileService())

	metrics.InitializeMetrics(server)
	reflection.Register(server)

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
