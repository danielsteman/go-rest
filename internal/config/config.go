package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `env:"PORT"`
	}
	Database struct {
		URL string `env:"DATABASE_URL"`
	}
}

func NewConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.SetDefault("server.port", "8080")

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
