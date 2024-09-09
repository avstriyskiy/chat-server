package config

import (
	"errors"
	"fmt"
	"os"
)

var (
	pgUser     = "POSTGRES_USER"
	pgPassword = "POSTGRES_PASSWORD" // nolint
	pgPort     = "POSTGRES_PORT"
	dbHost     = "DB_HOST"
	dbName     = "CHAT_SERVER_APP_DB"
)

// PostgresConfig struct
type PostgresConfig struct {
	PgUser     string
	PgPassword string
	PgPort     string
	DBName     string
	DBHost     string
}

// DBConfig interface
type DBConfig interface {
	DSN() string
}

// NewPostgresConfig creates new Postgresconfig instance
func NewPostgresConfig() (*PostgresConfig, error) {
	pgUser = os.Getenv(pgUser)
	pgPassword = os.Getenv(pgPassword)
	pgPort = os.Getenv(pgPort)
	dbHost = os.Getenv(dbHost)
	dbName = os.Getenv(dbName)

	if dbHost == "" || pgPassword == "" || pgPort == "" || pgUser == "" {
		return nil, errors.New("failed to initialize postgres config")
	}
	return &PostgresConfig{
		PgUser:     pgUser,
		PgPassword: pgPassword,
		PgPort:     pgPort,
		DBName:     dbName,
		DBHost:     dbHost,
	}, nil
}

// DSN Get Postgres DSN
func (cfg *PostgresConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.DBHost,
		cfg.PgPort,
		cfg.DBName,
		cfg.PgUser,
		cfg.PgPassword,
	)
}
