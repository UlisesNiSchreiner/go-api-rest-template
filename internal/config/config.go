package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Env   string
	HTTP  HTTP
	MySQL MySQL
}

type HTTP struct {
	Host string
	Port int
}

func (h HTTP) Addr() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

type MySQL struct {
	DSN string
}

func Load() (Config, error) {
	cfg := Config{
		Env: getenv("APP_ENV", "dev"),
		HTTP: HTTP{
			Host: getenv("HTTP_HOST", "0.0.0.0"),
			Port: getenvInt("HTTP_PORT", 8080),
		},
		MySQL: MySQL{
			DSN: getenv("MYSQL_DSN", ""),
		},
	}

	if cfg.MySQL.DSN == "" {
		return Config{}, errors.New("MYSQL_DSN is required")
	}

	return cfg, nil
}

func getenv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}

func getenvInt(key string, fallback int) int {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return fallback
	}
	return n
}
