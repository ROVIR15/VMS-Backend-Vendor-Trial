package delivery

import (
	"context"

	controller "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/controllers/rest"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/internal/vendors/interactor"
	echoServer "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/http/echo_server"

	"github.com/labstack/echo/v4"
)

type RestEndpointConfigurations struct {
	echoServer.EchoServer
	Group *echo.Group
}

func NewRestEndpointConfigurations(echo echoServer.EchoServer, base *echo.Group) *RestEndpointConfigurations {
	return &RestEndpointConfigurations{EchoServer: echo, Group: base.Group("/vendor")}
}

func (endpoint *RestEndpointConfigurations) RegisterRoutes(
	ctx context.Context,
	usecase *interactor.Usecase,
) {
	endpoint.RouteBuilder().RegisterRoutes(func(ec *echo.Echo) {
		// Register your rest endpoint here
		// example:
		// getExampleEndpoint := controller.NewRegisterUserController(endpoint.Group, endpoint.Logger(), authGuard, usecase.RegisterUserUsecase)
		// getExampleEndpoint.RegisterPath()

		getVendorEndpoint := controller.NewVendorController(endpoint.Group, endpoint.Logger(), usecase)
		getVendorEndpoint.RegisterPath()
	},
	)
}
