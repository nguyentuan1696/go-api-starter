package logger

import (
	"io"
	"os"

	"github.com/rs/zerolog"
	"github.com/samber/do/v2"
	"go-api-starter/core/config"
)

// Config holds the logger configuration.
type Config struct {
	Level   string
	Format  string
	Output  string
	NoColor bool
}

// NewLogger creates a new zerolog logger instance with dependency injection support
// This service is automatically registered with the do dependency injection container.
func NewLogger(i do.Injector) (*zerolog.Logger, error) {
	appConfig := do.MustInvoke[*config.Config](i)

	// Configure log level
	level, err := zerolog.ParseLevel(appConfig.Logger.Level)
	if err != nil {
		level = zerolog.InfoLevel
	}

	// Set global log level
	zerolog.SetGlobalLevel(level)

	// Configure output
	var output io.Writer
	if appConfig.Logger.Output == "stdout" || appConfig.Logger.Output == "" {
		output = zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    appConfig.Logger.NoColor,
			TimeFormat: "2006-01-02 15:04:05",
		}
	} else {
		//bearer:disable go_gosec_file_permissions_file_perm
		file, err := os.OpenFile(appConfig.Logger.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			// Fall back to stdout if file creation fails
			output = zerolog.ConsoleWriter{
				Out:        os.Stdout,
				NoColor:    true,
				TimeFormat: "2006-01-02 15:04:05",
			}
		} else {
			output = zerolog.MultiLevelWriter(zerolog.ConsoleWriter{
				Out:        file,
				NoColor:    true,
				TimeFormat: "2006-01-02 15:04:05",
			})
		}
	}

	// Create and configure logger
	logger := zerolog.New(output).With().Timestamp().Logger()

	return &logger, nil
}
