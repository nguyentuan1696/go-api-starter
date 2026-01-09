package config

import (
	"fmt"
	"strings"

	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config holds all application configuration
// This struct demonstrates how to structure configuration for dependency injection.
type Config struct {
	App        AppConfig        `mapstructure:"app"`
	Server     ServerConfig     `mapstructure:"http"`
	Postgresql PostgresqlConfig `mapstructure:"postgresql"`
	Redis      RedisConfig      `mapstructure:"redis"`
	Logger     LoggerConfig     `mapstructure:"logger"`
}

// ServerConfig holds HTTP http configuration.
type ServerConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

// DatabaseConfig holds PostgreSQL configuration.
type PostgresqlConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Database        string `mapstructure:"database"`
	SSLMode         string `mapstructure:"ssl_mode"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"`
}

type RedisConfig struct {
	Address  string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

// LoggerConfig holds logger configuration.
type LoggerConfig struct {
	Level   string `mapstructure:"level"`
	Format  string `mapstructure:"format"`
	Output  string `mapstructure:"output"`
	NoColor bool   `mapstructure:"no_color"`
}

// AppConfig holds application-specific configuration.
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
	Debug       bool   `mapstructure:"debug"`
}

// NewConfig creates a new configuration instance using viper
// This demonstrates configuration management with the samber/do library.
func NewConfig(i do.Injector) (*Config, error) {
	// Enable environment variable support
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	// Unmarshal configuration into struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

// SetCobraFlags adds command line flags to the cobra command
// This method demonstrates how services can provide functionality through DI.
func (cs *Config) SetCobraFlags(cmd *cobra.Command) {
	// Server flags
	_ = cmd.PersistentFlags().String("http.host", "localhost", "Server host")
	_ = cmd.PersistentFlags().Int("http.port", 8080, "Server port")
	_ = cmd.PersistentFlags().Int("http.read_timeout", 30, "Server read timeout in seconds")
	_ = cmd.PersistentFlags().Int("http.write_timeout", 30, "Server write timeout in seconds")

	// Database flags
	_ = cmd.PersistentFlags().String("database.host", "localhost", "Database host")
	_ = cmd.PersistentFlags().Int("database.port", 5432, "Database port")
	_ = cmd.PersistentFlags().String("database.user", "postgres", "Database user")
	_ = cmd.PersistentFlags().String("database.password", "postgres", "Database password")
	_ = cmd.PersistentFlags().String("database.database", "do_template_api", "Database name")
	_ = cmd.PersistentFlags().String("database.ssl_mode", "disable", "Database SSL mode")
	_ = cmd.PersistentFlags().Int("database.max_open_conns", 25, "Database max open connections")
	_ = cmd.PersistentFlags().Int("database.max_idle_conns", 25, "Database max idle connections")
	_ = cmd.PersistentFlags().Int("database.conn_max_lifetime", 300, "Database connection max lifetime in seconds")

	// Logger flags
	_ = cmd.PersistentFlags().String("logger.level", "info", "Log level")
	_ = cmd.PersistentFlags().String("logger.format", "console", "Log format")
	_ = cmd.PersistentFlags().String("logger.output", "stdout", "Log output")
	_ = cmd.PersistentFlags().Bool("logger.no_color", false, "Disable colored output")

	// App flags
	_ = cmd.PersistentFlags().String("app.name", "do-template-worker", "Application name")
	_ = cmd.PersistentFlags().String("app.version", "1.0.0", "Application version")
	_ = cmd.PersistentFlags().String("app.environment", "development", "Application environment")
	_ = cmd.PersistentFlags().Bool("app.debug", false, "Debug mode")

	// Bind all flags to viper for automatic configuration
	cs.bindFlagsToViper(cmd)
}

// bindFlagsToViper binds all cobra flags to viper.
func (cs *Config) bindFlagsToViper(cmd *cobra.Command) {
	// Server flags
	_ = viper.BindPFlag("http.host", cmd.PersistentFlags().Lookup("http.host"))
	_ = viper.BindPFlag("http.port", cmd.PersistentFlags().Lookup("http.port"))
	_ = viper.BindPFlag("http.read_timeout", cmd.PersistentFlags().Lookup("http.read_timeout"))
	_ = viper.BindPFlag("http.write_timeout", cmd.PersistentFlags().Lookup("http.write_timeout"))

	// Database flags
	_ = viper.BindPFlag("database.host", cmd.PersistentFlags().Lookup("database.host"))
	_ = viper.BindPFlag("database.port", cmd.PersistentFlags().Lookup("database.port"))
	_ = viper.BindPFlag("database.user", cmd.PersistentFlags().Lookup("database.user"))
	_ = viper.BindPFlag("database.password", cmd.PersistentFlags().Lookup("database.password"))
	_ = viper.BindPFlag("database.database", cmd.PersistentFlags().Lookup("database.database"))
	_ = viper.BindPFlag("database.ssl_mode", cmd.PersistentFlags().Lookup("database.ssl_mode"))
	_ = viper.BindPFlag("database.max_open_conns", cmd.PersistentFlags().Lookup("database.max_open_conns"))
	_ = viper.BindPFlag("database.max_idle_conns", cmd.PersistentFlags().Lookup("database.max_idle_conns"))
	_ = viper.BindPFlag("database.conn_max_lifetime", cmd.PersistentFlags().Lookup("database.conn_max_lifetime"))

	// Logger flags
	_ = viper.BindPFlag("logger.level", cmd.PersistentFlags().Lookup("logger.level"))
	_ = viper.BindPFlag("logger.format", cmd.PersistentFlags().Lookup("logger.format"))
	_ = viper.BindPFlag("logger.output", cmd.PersistentFlags().Lookup("logger.output"))
	_ = viper.BindPFlag("logger.no_color", cmd.PersistentFlags().Lookup("logger.no_color"))

	// App flags
	_ = viper.BindPFlag("app.name", cmd.PersistentFlags().Lookup("app.name"))
	_ = viper.BindPFlag("app.version", cmd.PersistentFlags().Lookup("app.version"))
	_ = viper.BindPFlag("app.environment", cmd.PersistentFlags().Lookup("app.environment"))
	_ = viper.BindPFlag("app.debug", cmd.PersistentFlags().Lookup("app.debug"))
}
