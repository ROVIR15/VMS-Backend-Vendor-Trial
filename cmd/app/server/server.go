package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/cmd/app/config"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/cmd/app/config/infrastructure"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/cmd/app/config/service"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/environment"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/echo_server/echo4"
	middleware "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/echo_server/middlewares/casbin"
	grpc "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/grpc"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
)

type Server struct {
	Logger logger.Logger
	Config *config.Config
}

func NewServer(logger logger.Logger, config *config.Config) *Server {
	return &Server{Logger: logger, Config: config}
}

func (s *Server) Run(e environment.Environment) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	echoServer := echo4.NewEcho4Server(s.Config.EchoOptions, s.Logger)

	casbinAdapter, err := middleware.NewCasbinAdapter("./pkg/casbin/rest/rbac_model.conf", "./pkg/casbin/rest/rbac_policy.csv")
	if err != nil {
		s.Logger.Errorf("[casbin] err: {%v}", err)
	}

	echoServer.AddMiddlewares(middleware.CasbinMiddleware(casbinAdapter))

	grpcServer := grpc.NewGrpcServer(s.Config.GrpcOptions, s.Logger)

	grpcClient, err := grpc.NewGrpcClient(&s.Config.GrpcClientOptions[0])
	if err != nil {
		s.Logger.Errorf("[grpc] err: {%v}", err)
		return err
	}

	infrastructureConfigurations, infrastructureCleanupFunc, err := infrastructure.NewInfrastructureConfigurator(s.Logger, s.Config).ConfigInfrastructures(ctx)
	if err != nil {
		s.Logger.Errorf("[infrastructure] err: {%v}", err)
		return err
	}
	defer infrastructureCleanupFunc()

	if err := service.NewServiceConfigurator(
		infrastructureConfigurations,
		echoServer,
		grpcServer,
		grpcClient,
	).ConfigureService(ctx); err != nil {
		s.Logger.Errorf("[service] err: {%v}", err)
		return err
	}

	// run http server
	if infrastructureConfigurations.Config.AppOptions.DeliveryType == string(config.HTTP) {
		s.Logger.Infof("Running http server")
		go func() {
			if err := echoServer.RunHttpServer(e, nil); err != http.ErrServerClosed {
				s.Logger.Errorf("[http] err: {%v}", err)
				stop()
			}

		}()
		s.Logger.Infof("%s http service is listening on PORT: {%d}", s.Config.AppOptions.GetMicroserviceName(), s.Config.EchoOptions.Port)

		// run grpc server
	} else if infrastructureConfigurations.Config.AppOptions.DeliveryType == string(config.GRPC) {
		s.Logger.Infof("Running grpc server")
		go func() {

			if err := grpcServer.RunGrpcServer(nil); err != nil {
				// do a fatal for going to OnStop process
				s.Logger.Errorf("[GrpcServer.RunGrpcServer] error in running server: {%v}", err)
			}
		}()
		s.Logger.Infof("%s grpc service is listening on PORT: {%d}", s.Config.AppOptions.GetMicroserviceName(), s.Config.GrpcOptions.Port)
	}

	// wait for signal to shutdown
	<-ctx.Done()

	s.Logger.Infof("%s is shutting down Http PORT: {%s}", s.Config.AppOptions.GetMicroserviceName(), s.Config.EchoOptions.Port)

	// stoping http server
	if infrastructureConfigurations.Config.AppOptions.DeliveryType == string(config.HTTP) {
		if err := echoServer.GracefulShutdown(ctx); err != nil {
			echoServer.Logger().Errorf("error shutting down echo server: %v", err)
		} else {
			echoServer.Logger().Info("echo server shutdown gracefully")
		}

		// stoping grpc server
	} else if infrastructureConfigurations.Config.AppOptions.DeliveryType == string(config.GRPC) {

		grpcServer.GracefulShutdown()
		s.Logger.Info("echo server shutdown gracefully")

	}

	s.Logger.Infof("%s server exited properly", s.Config.AppOptions.GetMicroserviceName())

	return nil
}
