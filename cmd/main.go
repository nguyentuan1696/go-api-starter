package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"go-api-starter/core/cli"
	"go-api-starter/core/config"
	"go-api-starter/core/http"
)

func main() {
	// Initialize the dependency injection injector
	// This is the core component of the samber/do library that manages all services
	injector := do.New(
		http.Package,
	)

	// Get services from dependency injection container
	appConfig := do.MustInvoke[*config.Config](injector)
	appLogger := do.MustInvoke[*zerolog.Logger](injector)
	cliService := do.MustInvoke[*cli.CLI](injector)

	// Start the application
	appLogger.Info().Str("app_name", appConfig.App.Name).
		Str("version", appConfig.App.Version).
		Str("environment", appConfig.App.Environment).
		Msg(fmt.Sprintf("Starting %s application", appConfig.App.Name))

	// Execute the CLI - this will handle all command parsing and execution
	if err := cliService.Execute(); err != nil {
	}

	_, _ = injector.ShutdownOnSignals()
}
