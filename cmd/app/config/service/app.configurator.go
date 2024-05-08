package service

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/cmd/app/config/infrastructure"
	echoServer "github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/http/echo_server"
	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/http/grpc"

	"github.com/labstack/echo/v4"
)

type ServiceConfigurator interface {
	ConfigureAuthModule() error
}

type serviceConfigurator struct {
	*infrastructure.InfrastructureConfigurations
	echoServer echoServer.EchoServer
	grpcServer grpc.GrpcServer
}

func NewServiceConfigurator(infra *infrastructure.InfrastructureConfigurations, echoServer echoServer.EchoServer, grpcServer grpc.GrpcServer) *serviceConfigurator {
	return &serviceConfigurator{echoServer: echoServer, grpcServer: grpcServer, InfrastructureConfigurations: infra}
}

func (c *serviceConfigurator) ConfigureService(ctx context.Context) error {

	// sharedUsecaseContainer := interactor.NewSharedUsecaseContainer()

	// define api base url
	// apiBaseURL := c.echoServer.GetEchoInstance().Group("/api")

	// ModuleConfigurator := Module.NewModuleConfigurator(c.InfrastructureConfigurations, c.echoServer, apiBaseURL, c.grpcServer)
	// if err := ModuleConfigurator.ConfigureModule(ctx, sharedUsecaseContainer); err != nil {
	// 	return errors.Wrap(err, "error while configuring module")
	// }

	c.echoServer.RouteBuilder().
		RegisterRoutes(func(e *echo.Echo) {
			e.GET("", func(ec echo.Context) error {
				return ec.String(
					http.StatusOK,
					fmt.Sprintf(
						"%s is running...",
						c.Config.AppOptions.GetMicroserviceNameUpper(),
					),
				)
			})
		})

	return nil
}
