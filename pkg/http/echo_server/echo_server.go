package echoServer

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/environment"
	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/http/echo_server/config"
	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/logger"

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
