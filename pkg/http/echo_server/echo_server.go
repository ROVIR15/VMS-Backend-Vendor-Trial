package echoServer

import (
	"context"

	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/environment"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/http/echo_server/config"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/logger"

	"github.com/labstack/echo/v4"
)

type EchoServer interface {
	RunHttpServer(e environment.Environment, configEcho ...func(echo *echo.Echo)) error
	GracefulShutdown(ctx context.Context) error
	ApplyVersioningFromHeader()
	GetEchoInstance() *echo.Echo
	Logger() logger.Logger
	Options() *config.EchoOptions
	SetupDefaultMiddlewares()
	RouteBuilder() *RouteBuilder
	AddMiddlewares(middlewares ...echo.MiddlewareFunc)
}
