package server

import (
	"go-api-starter/pkg/config"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
)

type HTTPServer struct {
	config *config.Config  `do:""`
	logger *zerolog.Logger `do:""`
	Server *http.Server
	Engine *echo.Echo
}

func NewHTTPServer(injector do.Injector) (*HTTPServer, error) {
	server := do.MustInvokeStruct[*HTTPServer](injector)

	server.Engine = echo.New()

	server.Engine.Use(middleware.CORS())

	server.Engine.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogLatency: true,
		LogMethod:  true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			server.logger.Info().
				Str("method", values.Method).
				Str("uri", values.URI).
				Int("status", values.Status).
				Dur("latency", values.Latency).
				Msg("Request")
			return nil
		},
	}))

	// Configure HTTP server using config from dependency injection
	server.Server = &http.Server{
		Addr:         server.config.Server.Host + ":" + strconv.Itoa(server.config.Server.Port),
		Handler:      server.Engine,
		ReadTimeout:  time.Duration(server.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(server.config.Server.WriteTimeout) * time.Second,
		IdleTimeout:  60 * time.Second, // Default idle timeout
	}

	return server, nil

}

func (s *HTTPServer) Start() error {
	s.logger.Info().
		Str("host", s.config.Server.Host).
		Int("port", s.config.Server.Port).
		Msg("Starting HTTP server")

	return s.Server.ListenAndServe()
}
