package server

import (
	"github.com/Elfsilon/noty-backend/internal/app/store"
)

// Config ...
type Config struct {
	Addr     string `toml:"addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Addr:     ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
