package config

import (
	"fmt"
	"strings"

	"github.com/samber/do/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig     `mapstructure:"server"`
	Redis      RedisConfig      `mapstructure:"redis"`
	Postgresql PostgresqlConfig `mapstructure:"postgresql"`
	Logger     LoggerConfig     `mapstructure:"logger"`
	App        AppConfig        `mapstructure:"app"`
	Minio      MinioConfig      `mapstructure:"minio"`
}

type ServerConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

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

type LoggerConfig struct {
	Level   string `mapstructure:"level"`
	Format  string `mapstructure:"format"`
	Output  string `mapstructure:"output"`
	NoColor bool   `mapstructure:"no_color"`
}

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
	Debug       bool   `mapstructure:"debug"`
}

type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	UseSSL          bool   `mapstructure:"use_ssl"`
}

func NewConfig(i do.Injector) (*Config, error) {
	// Enable environment variable support
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set config file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		// It's okay if config file doesn't exist, we fallback to flags/env
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	// Unmarshal configuration into struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
}

func (cs *Config) SetCobraFlags(cmd *cobra.Command) {
	// Server flags
	_ = cmd.PersistentFlags().String("server.host", "localhost", "Server host")
	_ = cmd.PersistentFlags().Int("server.port", 8080, "Server port")
	_ = cmd.PersistentFlags().Int("server.read_timeout", 30, "Server read timeout in seconds")
	_ = cmd.PersistentFlags().Int("server.write_timeout", 30, "Server write timeout in seconds")

	// Redis flags
	_ = cmd.PersistentFlags().String("redis.host", "localhost", "Redis host")
	_ = cmd.PersistentFlags().Int("redis.port", 6379, "Redis port")
	_ = cmd.PersistentFlags().String("redis.password", "", "Redis password")
	_ = cmd.PersistentFlags().Int("redis.db", 0, "Redis database")

	// Database flags
	_ = cmd.PersistentFlags().String("postgresql.host", "localhost", "Database host")
	_ = cmd.PersistentFlags().Int("postgresql.port", 5432, "Database port")
	_ = cmd.PersistentFlags().String("postgresql.user", "postgres", "Database user")
	_ = cmd.PersistentFlags().String("postgresql.password", "postgres", "Database password")
	_ = cmd.PersistentFlags().String("postgresql.database", "do_template_api", "Database name")
	_ = cmd.PersistentFlags().String("postgresql.ssl_mode", "disable", "Database SSL mode")
	_ = cmd.PersistentFlags().Int("postgresql.max_open_conns", 25, "Database max open connections")
	_ = cmd.PersistentFlags().Int("postgresql.max_idle_conns", 25, "Database max idle connections")
	_ = cmd.PersistentFlags().Int("postgresql.conn_max_lifetime", 300, "Database connection max lifetime in seconds")

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

	// Minio flags
	_ = cmd.PersistentFlags().String("minio.endpoint", "localhost", "Minio endpoint")
	_ = cmd.PersistentFlags().String("minio.access_key_id", "minioadmin", "Minio access key ID")
	_ = cmd.PersistentFlags().String("minio.secret_access_key", "minioadmin", "Minio secret access key")
	_ = cmd.PersistentFlags().Bool("minio.use_ssl", false, "Minio use SSL")

	// Bind all flags to viper for automatic configuration
	cs.bindFlagsToViper(cmd)
}

// bindFlagsToViper binds all cobra flags to viper.
func (cs *Config) bindFlagsToViper(cmd *cobra.Command) {
	// Server flags
	_ = viper.BindPFlag("server.host", cmd.PersistentFlags().Lookup("server.host"))
	_ = viper.BindPFlag("server.port", cmd.PersistentFlags().Lookup("server.port"))
	_ = viper.BindPFlag("server.read_timeout", cmd.PersistentFlags().Lookup("server.read_timeout"))
	_ = viper.BindPFlag("server.write_timeout", cmd.PersistentFlags().Lookup("server.write_timeout"))

	// Redis flags
	_ = viper.BindPFlag("redis.host", cmd.PersistentFlags().Lookup("redis.host"))
	_ = viper.BindPFlag("redis.port", cmd.PersistentFlags().Lookup("redis.port"))
	_ = viper.BindPFlag("redis.password", cmd.PersistentFlags().Lookup("redis.password"))
	_ = viper.BindPFlag("redis.db", cmd.PersistentFlags().Lookup("redis.db"))

	// Database flags
	_ = viper.BindPFlag("postgresql.host", cmd.PersistentFlags().Lookup("postgresql.host"))
	_ = viper.BindPFlag("postgresql.port", cmd.PersistentFlags().Lookup("postgresql.port"))
	_ = viper.BindPFlag("postgresql.user", cmd.PersistentFlags().Lookup("postgresql.user"))
	_ = viper.BindPFlag("postgresql.password", cmd.PersistentFlags().Lookup("postgresql.password"))
	_ = viper.BindPFlag("postgresql.database", cmd.PersistentFlags().Lookup("postgresql.database"))
	_ = viper.BindPFlag("postgresql.ssl_mode", cmd.PersistentFlags().Lookup("postgresql.ssl_mode"))
	_ = viper.BindPFlag("postgresql.max_open_conns", cmd.PersistentFlags().Lookup("postgresql.max_open_conns"))
	_ = viper.BindPFlag("postgresql.max_idle_conns", cmd.PersistentFlags().Lookup("postgresql.max_idle_conns"))
	_ = viper.BindPFlag("postgresql.conn_max_lifetime", cmd.PersistentFlags().Lookup("postgresql.conn_max_lifetime"))

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

	// Minio flags
	_ = viper.BindPFlag("minio.endpoint", cmd.PersistentFlags().Lookup("minio.endpoint"))
	_ = viper.BindPFlag("minio.access_key_id", cmd.PersistentFlags().Lookup("minio.access_key_id"))
	_ = viper.BindPFlag("minio.secret_access_key", cmd.PersistentFlags().Lookup("minio.secret_access_key"))
	_ = viper.BindPFlag("minio.use_ssl", cmd.PersistentFlags().Lookup("minio.use_ssl"))
}
