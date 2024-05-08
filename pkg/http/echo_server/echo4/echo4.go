package echo4

import (
	"context"
	"fmt"
	"strings"

	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/constants"
	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/environment"
	echoServer "github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/http/echo_server"
	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/http/echo_server/config"
	customHandlers "github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/http/echo_server/handlers"
	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/http/echo_server/middlewares/log"
	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/logger"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

type echo4Server struct {
	echo         *echo.Echo
	options      *config.EchoOptions
	log          logger.Logger
	routeBuilder *echoServer.RouteBuilder
}

// echo4Server implements EchoServer
var _ echoServer.EchoServer = (*echo4Server)(nil)

func NewEcho4Server(
	options *config.EchoOptions,
	logger logger.Logger,
) *echo4Server {
	e := echo.New()
	e.HideBanner = true

	return &echo4Server{
		echo:         e,
		options:      options,
		log:          logger,
		routeBuilder: echoServer.NewRouteBuilder(e),
	}
}

func (s *echo4Server) RunHttpServer(
	e environment.Environment,
	configEcho ...func(echo *echo.Echo),
) error {
	s.echo.Server.ReadTimeout = constants.ReadTimeout
	s.echo.Server.WriteTimeout = constants.WriteTimeout
	s.echo.Server.MaxHeaderBytes = constants.MaxHeaderBytes

	if len(configEcho) > 0 {
		ehcoFunc := configEcho[0]
		if ehcoFunc != nil {
			configEcho[0](s.echo)
		}
	}

	// setup default middlewares
	s.SetupDefaultMiddlewares()

	if e.IsStaging() || e.IsProduction() {
		s.log.Infow("[ECHO] use tls", logger.Fields{"IsTLS": true})
		s.echo.AutoTLSManager.Cache = autocert.DirCache(s.options.TLSCacheDir)
		s.echo.AutoTLSManager.HostPolicy = autocert.HostWhitelist(s.options.TLSHosts...)

		s.echo.Use(middleware.HTTPSRedirect())

		return s.echo.StartAutoTLS(s.options.Port)
	} else {
		// https://echo.labstack.com/guide/http_server/
		return s.echo.Start(s.options.Port)
	}
}

func (s *echo4Server) Logger() logger.Logger {
	return s.log
}

func (s *echo4Server) Options() *config.EchoOptions {
	return s.options
}

func (s *echo4Server) RouteBuilder() *echoServer.RouteBuilder {
	return s.routeBuilder
}

func (s *echo4Server) AddMiddlewares(middlewares ...echo.MiddlewareFunc) {
	if len(middlewares) > 0 {
		s.echo.Use(middlewares...)
	}
}

func (s *echo4Server) GracefulShutdown(ctx context.Context) error {
	err := s.echo.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (s *echo4Server) SetupDefaultMiddlewares() {

	// apply versioning from header
	// s.ApplyVersioningFromHeader()

	// set error handler
	s.echo.HTTPErrorHandler = func(err error, c echo.Context) {
		// bypass notfound favicon endpoint and its error
		fmt.Println(c.Request().URL.Path)
		if c.Request().URL.Path == "/favicon.ico" {
			return
		}
		customHandlers.ProblemHandlerFunc(err, c, s.log)
	}

	// log errors and information
	s.echo.Use(log.EchoLoggerMiddleware(s.log))

	// s.echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
	// 	LogContentLength: true,
	// 	LogLatency:       true,
	// 	LogError:         false,
	// 	LogMethod:        true,
	// 	LogRequestID:     true,
	// 	LogURI:           true,
	// 	LogResponseSize:  true,
	// 	LogURIPath:       true,
	// 	LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
	// 		s.log.Infow(
	// 			fmt.Sprintf("[Request Middleware] REQUEST: uri: %v, status: %v", v.URI, v.Status),
	// 			logger.Fields{"URI": v.URI, "Status": v.Status},
	// 		)
	// 		return nil
	// 	},
	// }))

	s.echo.Use(middleware.BodyLimit(constants.BodyLimit))
	s.echo.Use(middleware.RequestID())
	s.echo.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: constants.GzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
}

func (s *echo4Server) ApplyVersioningFromHeader() {
	s.echo.Pre(apiVersion)
}

func (s *echo4Server) GetEchoInstance() *echo.Echo {
	return s.echo
}

// APIVersion Header Based Versioning
func apiVersion(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()
		headers := req.Header

		apiVersion := headers.Get("version")
		if apiVersion == "" {
			apiVersion = "v1"
		}

		route := strings.Split(req.URL.Path, "/")
		if route[1] == "api" {
			req.URL.Path = fmt.Sprintf("/api/%s/%s", apiVersion, strings.Join(route[2:], "/"))
		}

		return next(c)
	}
}
