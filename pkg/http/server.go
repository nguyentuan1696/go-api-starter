package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/do/v2"
	"go-api-starter/core/config"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	config *config.Config `do:""`
	server *http.Server
	engine *echo.Echo
}

func NewHTTPServer(injector do.Injector) (*Server, error) {
	server := do.MustInvokeStruct[*Server](injector)

	server.engine = echo.New()
	server.engine.Use(middleware.Logger())
	server.engine.Use(middleware.Recover())

	server.server = &http.Server{
		Addr:         server.config.Server.Host + ":" + strconv.Itoa(server.config.Server.Port),
		Handler:      server.engine,
		ReadTimeout:  time.Duration(server.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(server.config.Server.WriteTimeout) * time.Second,
		IdleTimeout:  60 * time.Second, // Default idle timeout
	}

	return server, nil

}

// Start starts the HTTP server
// This demonstrates how to start the server with proper error handling.
func (s *Server) Start() error {

	return s.server.ListenAndServe()
}

func (s *Server) ShutdownWithContext(ctx context.Context) error {
	return s.engine.Shutdown(ctx)

}
