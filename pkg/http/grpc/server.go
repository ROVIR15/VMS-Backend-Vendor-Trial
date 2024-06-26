package grpc

import (
	"net"
	"time"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/grpc/config"
	grpcError "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/grpc/interceptors/grpc_error"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/utils"

	"emperror.dev/errors"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	googleGrpc "google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

const (
	maxConnectionIdle = 5
	gRPCTimeout       = 15
	maxConnectionAge  = 5
	gRPCTime          = 10
)

type GrpcServer interface {
	RunGrpcServer(configGrpc ...func(grpcServer *googleGrpc.Server)) error
	GracefulShutdown()
	GetCurrentGrpcServer() *googleGrpc.Server
	GrpcServiceBuilder() *GrpcServiceBuilder
	Logger() logger.Logger
}

type grpcServer struct {
	server         *googleGrpc.Server
	config         *config.GrpcOptions
	log            logger.Logger
	serviceName    string
	serviceBuilder *GrpcServiceBuilder
}

func NewGrpcServer(
	config *config.GrpcOptions,
	logger logger.Logger,
) GrpcServer {
	unaryServerInterceptors := []googleGrpc.UnaryServerInterceptor{
		grpcError.UnaryServerInterceptor(),
		grpcCtxTags.UnaryServerInterceptor(),
		grpcRecovery.UnaryServerInterceptor(),
	}
	streamServerInterceptors := []googleGrpc.StreamServerInterceptor{
		grpcError.StreamServerInterceptor(),
	}

	s := googleGrpc.NewServer(
		googleGrpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: maxConnectionIdle * time.Minute,
			Timeout:           gRPCTimeout * time.Second,
			MaxConnectionAge:  maxConnectionAge * time.Minute,
			Time:              gRPCTime * time.Minute,
		}),
		// https://github.com/open-telemetry/opentelemetry-go-contrib/tree/00b796d0cdc204fa5d864ec690b2ee9656bb5cfc/instrumentation/google.golang.org/grpc/otelgrpc
		// github.com/grpc-ecosystem/go-grpc-middleware
		googleGrpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			streamServerInterceptors...,
		)),
		googleGrpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			unaryServerInterceptors...,
		)),
	)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthServer)
	healthServer.SetServingStatus(config.Name, grpc_health_v1.HealthCheckResponse_SERVING)

	return &grpcServer{
		server:         s,
		config:         config,
		log:            logger,
		serviceName:    config.Name,
		serviceBuilder: NewGrpcServiceBuilder(s),
	}
}

func (s *grpcServer) RunGrpcServer(
	configGrpc ...func(grpcServer *googleGrpc.Server),
) error {
	l, err := net.Listen("tcp", utils.ColonPort(s.config.Port))
	if err != nil {
		return errors.WrapIf(err, "net.Listen")
	}

	if len(configGrpc) > 0 {
		grpcFunc := configGrpc[0]
		if grpcFunc != nil {
			grpcFunc(s.server)
		}
	}

	if s.config.Development {
		reflection.Register(s.server)
	}

	err = s.server.Serve(l)

	return err
}

func (s *grpcServer) GrpcServiceBuilder() *GrpcServiceBuilder {
	return s.serviceBuilder
}

func (s *grpcServer) GetCurrentGrpcServer() *googleGrpc.Server {
	return s.server
}

func (s *grpcServer) GracefulShutdown() {
	s.server.Stop()
	s.server.GracefulStop()
}

func (s *grpcServer) Logger() logger.Logger {
	return s.log
}
