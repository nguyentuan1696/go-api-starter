package cli

import (
	"context"
	"go-api-starter/pkg/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	authHTTPRouter "go-api-starter/modules/auth/router/http"
	serverService "go-api-starter/pkg/server"
)

type CLI struct {
	config      *config.Config `do:""`
	injector    do.Injector
	rootCommand *cobra.Command
}

// NewCLI creates a new CLI service with dependency injection support.
func NewCLI(i do.Injector) (*CLI, error) {
	cli := &CLI{
		config:   do.MustInvoke[*config.Config](i),
		injector: i,
	}

	// Create the root command
	cli.rootCommand = &cobra.Command{
		Use:     cli.config.App.Name,
		Short:   "AnawimEnglish API",
		Version: cli.config.App.Version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// Re-unmarshal config to pick up flag values
			return viper.Unmarshal(cli.config)
		},
	}

	// Add persistent flags using dependency injection
	cli.setupPersistentFlags()

	// Add commands
	cli.setupCommands()

	return cli, nil
}

// setupPersistentFlags adds global flags to the CLI.
func (cli *CLI) setupPersistentFlags() {
	// Use the config service to set up all configuration flags
	// This demonstrates dependency injection for configuration management
	cli.config.SetCobraFlags(cli.rootCommand)
}

// setupCommands adds subcommands to the CLI.
func (cli *CLI) setupCommands() {
	// Add serve command
	cli.rootCommand.AddCommand(cli.newServeCommand())

}

// newServeCommand creates the serve command.
func (cli *CLI) newServeCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Start the Go API Starter service",
		Run: func(cmd *cobra.Command, args []string) {
			// Get the HTTP server from the dependency injection container
			httpServer := do.MustInvoke[*serverService.HTTPServer](cli.injector)
			logger := do.MustInvoke[*zerolog.Logger](cli.injector)

			// Register routes
			auth := do.MustInvoke[*authHTTPRouter.AuthHTTPRouter](cli.injector)
			auth.Register(httpServer.Engine)

			// Setup graceful shutdown
			_, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Setup signal handling
			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

			// Start server in goroutine
			go func() {
				logger.Info().Msg("Starting HTTP server...")
				if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {
					logger.Fatal().Err(err).Msg("Failed to start HTTP server")
				}
			}()

			// Wait for signal
			<-sigChan
			logger.Info().Msg("Shutting down...")
		},
	}
}

// RootCommand returns the root cobra command.
func (cli *CLI) RootCommand() *cobra.Command {
	return cli.rootCommand
}

// Execute executes the CLI with the given arguments.
func (cli *CLI) Execute() error {
	return cli.rootCommand.Execute()
}

// AddCommand adds a new command to the CLI.
func (cli *CLI) AddCommand(command *cobra.Command) {
	cli.rootCommand.AddCommand(command)
}
