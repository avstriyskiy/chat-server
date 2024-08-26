package config

import (
	"errors"
	"os"
)

var (
	authPort       = "AUTH_APP_PORT"
	chatServerPort = "CHAT_SERVER_APP_PORT"
)

type appConfig struct {
	authPort       string
	chatServerPort string
}

// AppConfig interface
type AppConfig interface {
	GRPCPort() string
}

// NewAppConfig generator
func NewAppConfig() (*appConfig, error) {
	authPort = os.Getenv(authPort)
	chatServerPort = os.Getenv(chatServerPort)

	if authPort == "" {
		return nil, errors.New("failed to initialize auth config")
	}
	return &appConfig{
		authPort:       authPort,
		chatServerPort: chatServerPort,
	}, nil
}

func (cfg *appConfig) GRPCPort() string {
	return cfg.chatServerPort
}
