package service

import (
	"context"

	"emperror.dev/errors"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/cmd/app/config/infrastructure"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/shared/interactor"
	vendorModel "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/config"

	echoServer "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/echo_server"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/grpc"
)

type ServiceConfigurator interface {
	ConfigureAuthModule() error
}

type serviceConfigurator struct {
	*infrastructure.InfrastructureConfigurations
	echoServer echoServer.EchoServer
	grpcServer grpc.GrpcServer
	grpcClient grpc.GrpcClient
}

func NewServiceConfigurator(infra *infrastructure.InfrastructureConfigurations, echoServer echoServer.EchoServer, grpcServer grpc.GrpcServer, grpcClient grpc.GrpcClient) *serviceConfigurator {
	return &serviceConfigurator{echoServer: echoServer, grpcServer: grpcServer, grpcClient: grpcClient, InfrastructureConfigurations: infra}
}

func (c *serviceConfigurator) ConfigureService(ctx context.Context) error {

	sharedUsecaseContainer := interactor.NewSharedUsecaseContainer()

	// define api base url
	apiBaseURL := c.echoServer.GetEchoInstance().Group("/api")
	ModuleConfigurator := vendorModel.NewModuleConfigurator(c.InfrastructureConfigurations, c.echoServer, apiBaseURL, c.grpcServer)
	if err := ModuleConfigurator.ConfigureModule(ctx, sharedUsecaseContainer); err != nil {
		return errors.Wrap(err, "error while configuring module")
	}

	// c.echoServer.RouteBuilder().
	// 	RegisterRoutes(func(e *echo.Echo) {
	// 		e.GET("", func(ec echo.Context) error {
	// 			return ec.String(
	// 				http.StatusOK,
	// 				fmt.Sprintf(
	// 					"%s is running...",
	// 					c.Config.AppOptions.GetMicroserviceNameUpper(),
	// 				),
	// 			)
	// 		})
	// 	})

	return nil
}
