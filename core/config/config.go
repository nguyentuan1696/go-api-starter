package config

import (
	"fmt"
	"github.com/samber/do/v2"
	"github.com/spf13/viper"
	"strings"
)

type Environment string

const (
	DevEnvironment  Environment = "dev"
	ProdEnvironment Environment = "prod"
)

type Config struct {
	Environment Environment
	Server      ServerConfig
	Database    DatabaseConfig
	JWT         JWTConfig
	Redis       RedisConfig
}

type ServerConfig struct {
	Port    int    `mapstructure:"port"`
	Host    string `mapstructure:"host"`
	BaseURL string `mapstructure:"base_url"`
}

type DatabaseConfig struct {
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

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func NewConfig(i do.Injector) (*Config, error) {
	v := viper.New()

	v.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	// 3. Unmarshal
	var config Config
	if err := v.Unmarshal(config); err != nil {
		err = fmt.Errorf("unable to decode config into struct: %w", err)
		return nil, nil
	}
	return &config, nil
}
