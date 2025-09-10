package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// ----------------------------------------------------------------------------
// Environment
// ----------------------------------------------------------------------------

type Environment string

const (
	DevEnvironment  Environment = "dev"
	ProdEnvironment Environment = "prod"
)

// ----------------------------------------------------------------------------
// Structs cấu hình
// ----------------------------------------------------------------------------

type ServerConfig struct {
	Port    int    `mapstructure:"port"`
	Host    string `mapstructure:"host"`
	BaseURL string `mapstructure:"base_url"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

type SMTPConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	FromName string `mapstructure:"from_name"`
	From     string `mapstructure:"from"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type R2Config struct {
	Endpoint        string `mapstructure:"endpoint"`
	PublicEndpoint  string `mapstructure:"public_endpoint"`
	Bucket          string `mapstructure:"bucket"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	Region          string `mapstructure:"region"`
}

type Config struct {
	Environment Environment
	Server      ServerConfig
	Database    DatabaseConfig
	JWT         JWTConfig
	SMTP        SMTPConfig
	Redis       RedisConfig `mapstructure:"redis"`
	R2          R2Config    `mapstructure:"r2"`
}

// ----------------------------------------------------------------------------
// Singleton
// ----------------------------------------------------------------------------

var (
	instance *Config
	once     sync.Once
)

func Get() *Config {
	if instance == nil {
		return &Config{
			Environment: DevEnvironment,
			Server:      ServerConfig{Port: 7070, Host: "localhost"},
			Database:    DatabaseConfig{Port: 5432},
		}
	}
	return instance
}

func GetSafe() (*Config, bool) { return instance, instance != nil }

// ----------------------------------------------------------------------------
// Validation helpers
// ----------------------------------------------------------------------------

func (c *Config) Validate() error {
	var errors []string

	if c.Database.Host == "" {
		errors = append(errors, "database host is required")
	}
	if c.Database.Port <= 0 || c.Database.Port > 65535 {
		errors = append(errors, "database port must be between 1 and 65535")
	}
	if c.Database.User == "" {
		errors = append(errors, "database user is required")
	}
	if c.Database.DBName == "" {
		errors = append(errors, "database name is required")
	}
	if c.JWT.Secret == "" {
		errors = append(errors, "JWT secret is required")
	}

	if c.R2.AccessKeyID != "" || c.R2.SecretAccessKey != "" || c.R2.Endpoint != "" {
		if c.R2.AccessKeyID == "" {
			errors = append(errors, "R2 access key ID is required when R2 is configured")
		}
		if c.R2.SecretAccessKey == "" {
			errors = append(errors, "R2 secret access key is required when R2 is configured")
		}
		if c.R2.Endpoint == "" {
			errors = append(errors, "R2 endpoint is required when R2 is configured")
		}
		if c.R2.Region == "" {
			errors = append(errors, "R2 region is required when R2 is configured")
		}
		if c.R2.Bucket == "" {
			errors = append(errors, "R2 bucket is required when R2 is configured")
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("configuration validation failed:\n- %s", strings.Join(errors, "\n- "))
	}
	return nil
}

// ----------------------------------------------------------------------------
// Internal helpers
// ----------------------------------------------------------------------------

func isRunningInDocker() bool {
	if _, ok := os.LookupEnv("DOCKER_CONTAINER"); ok {
		return true
	}
	if data, err := os.ReadFile("/proc/1/cgroup"); err == nil &&
		strings.Contains(string(data), "docker") {
		return true
	}
	return false
}

// ----------------------------------------------------------------------------
// Init
// ----------------------------------------------------------------------------

// Trong hàm Init, thêm bind cho SMTP config
func Init(env Environment) error {
	var err error
	once.Do(func() {
		// 1. Load .env nếu chạy local / ngoài Docker
		if !isRunningInDocker() {
			if _, err = os.Stat(".env"); err == nil {
				_ = godotenv.Load(".env")
			}
		}

		// 2. Khởi tạo Viper – chỉ lấy từ env
		v := viper.New()
		v.SetDefault("environment", string(env))
		v.SetDefault("server.port", 7070)
		v.SetDefault("server.host", "0.0.0.0")
		// v.SetDefault("database.port", 5432)
		// v.SetDefault("smtp.port", 587)
		// v.SetDefault("smtp.host", "smtp.gmail.com")
		// v.SetDefault("jwt.secret", "12345678901234567890123456789012")

		v.AutomaticEnv()
		v.SetEnvPrefix("APP")
		v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		// Bind all needed envs
		// Server configuration
		v.BindEnv("server.host", "APP_SERVER_HOST")
		v.BindEnv("server.base_url", "APP_SERVER_BASE_URL")

		// Database configuration
		v.BindEnv("database.host", "APP_DATABASE_HOST")
		v.BindEnv("database.port", "APP_DATABASE_PORT")
		v.BindEnv("database.user", "APP_DATABASE_USER")
		v.BindEnv("database.password", "APP_DATABASE_PASSWORD")
		v.BindEnv("database.dbname", "APP_DATABASE_DBNAME")

		// SMTP configuration
		v.BindEnv("smtp.host", "APP_SMTP_HOST")
		v.BindEnv("smtp.port", "APP_SMTP_PORT")
		v.BindEnv("smtp.username", "APP_SMTP_USERNAME")
		v.BindEnv("smtp.password", "APP_SMTP_PASSWORD")
		v.BindEnv("smtp.from", "APP_SMTP_FROM")
		v.BindEnv("smtp.from_name", "APP_SMTP_FROM_NAME")

		// Redis configuration
		v.BindEnv("redis.address", "APP_REDIS_ADDRESS")
		v.BindEnv("redis.password", "APP_REDIS_PASSWORD")
		v.BindEnv("redis.db", "APP_REDIS_DB")

		// R2 configuration
		v.BindEnv("r2.endpoint", "APP_R2_ENDPOINT")
		v.BindEnv("r2.public_endpoint", "APP_R2_PUBLIC_ENDPOINT")
		v.BindEnv("r2.bucket", "APP_R2_BUCKET")
		v.BindEnv("r2.access_key_id", "APP_R2_ACCESS_KEY_ID")
		v.BindEnv("r2.secret_access_key", "APP_R2_SECRET_ACCESS_KEY")
		v.BindEnv("r2.region", "APP_R2_REGION")

		// JWT configuration
		v.BindEnv("jwt.secret", "APP_JWT_SECRET")

		// 3. Unmarshal
		instance = &Config{}
		if err = v.Unmarshal(instance); err != nil {
			err = fmt.Errorf("unable to decode config into struct: %w", err)
			return
		}
		instance.Environment = env
	})
	return err
}
