package config

import (
	"fmt"
	"time"
)

type Config struct {
	Server    ServerConfig    `mapstructure:"server"`
	Database  DatabaseConfig  `mapstructure:"database"`
	Logger    LoggerConfig    `mapstructure:"logger"`
	Migration MigrationConfig `mapstructure:"migrations"`
}

type ServerConfig struct {
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	IdleTimeout  time.Duration `mapstructure:"idle_timeout"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSlMode  string `mapstructure:"ssl_mode"`
}

type LoggerConfig struct {
	Level string `mapstructure:"level"`
}

type MigrationConfig struct {
	Dir string `mapstructure:"dir"`
}

func (cfg *DatabaseConfig) ConStr() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SSlMode,
	)
}

func (cfg *ServerConfig) PortStr() string {
	return fmt.Sprintf(":%d", cfg.Port)
}
