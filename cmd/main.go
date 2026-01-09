package main

import (
	"go-api-starter/modules/auth"
	"go-api-starter/pkg"
	"go-api-starter/pkg/cli"
	"go-api-starter/pkg/config"
	"go-api-starter/pkg/server"

	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
)

func main() {

	injector := do.New(
		pkg.BasePackage,
		server.Package,
		auth.Package,
	)
	defer injector.Shutdown()

	appConfig := do.MustInvoke[*config.Config](injector)
	appLogger := do.MustInvoke[*zerolog.Logger](injector)
	cliService := do.MustInvoke[*cli.CLI](injector)

	// Start the application
	appLogger.Info().Str("app_name", appConfig.App.Name).
		Str("version", appConfig.App.Version).
		Str("environment", appConfig.App.Environment).
		Msg("Starting AnawimEnglish application")

	// Execute the CLI - this will handle all command parsing and execution
	if err := cliService.Execute(); err != nil {
		appLogger.Fatal().Err(err).Msg("Failed to execute API")
	}
}
