package config

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/cmd/app/config/infrastructure"
	sharedInteractor "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/shared/interactor"
	delivery "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/config/delivery/rest"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/config/mapping"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/interactor"

	repository "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/repositories/impl"

	echoServer "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/echo_server"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/grpc"
	"github.com/labstack/echo/v4"
)

type moduleConfigurator struct {
	*infrastructure.InfrastructureConfigurations
	echoServer echoServer.EchoServer
	apiBaseURL *echo.Group
	grpcServer grpc.GrpcServer
}

func NewModuleConfigurator(infrastructure *infrastructure.InfrastructureConfigurations, echoServer echoServer.EchoServer, apiBaseURL *echo.Group, grpcServer grpc.GrpcServer) *moduleConfigurator {
	return &moduleConfigurator{InfrastructureConfigurations: infrastructure, echoServer: echoServer, apiBaseURL: apiBaseURL, grpcServer: grpcServer}
}

func (c *moduleConfigurator) ConfigureModule(ctx context.Context, sharedUsecaseContainer sharedInteractor.SharedUsecaseContainer) error {
	if err := mapping.ConfigureMappings(); err != nil {
		return err
	}

	// eventPublisher := pubsub.NewEventPublisher(c.Logger, c.CloudEventsClient)

	vendorRepository := repository.NewSQLVendorRepositoryConn(c.DB)

	Usecases := interactor.NewUseCase(c.Logger, vendorRepository)

	Endpoint := delivery.NewRestEndpointConfigurations(c.echoServer, c.apiBaseURL)
	Endpoint.RegisterRoutes(ctx, Usecases)

	// grpcEndpoint := grpcDelivery.NewGrpcEndpointConfigurations(c.grpcServer)
	// grpcEndpoint.RegisterRoutes(ctx, Usecases)

	return nil
}
