package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal/env"
)

const (
	pgHostEnvKey     = "PGHOST"
	pgPortEnvKey     = "PGPORT"
	pgNameEnvKey     = "PGDATABASE"
	pgUsernameEnvKey = "PGUSER"
	pgPasswordEnvKey = "PGPASSWORD"
)

var (
	errBuildingPostgresConfig = func(err error) error {
		return fmt.Errorf("error when building postgres config: %w", err)
	}
)

type PostgresConfig interface {
	Source() string
}

type postgresConfig struct {
	host     string
	port     string
	name     string
	username string
	password string
}

func (c postgresConfig) Source() string {
	return fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=require", c.port, c.host, c.username, c.password, c.name)
}

func NewPostgresConfig() (PostgresConfig, error) {

	host, err := env.MustGetEnvironmentVariable(pgHostEnvKey)
	if err != nil {
		return postgresConfig{}, errBuildingPostgresConfig(err)
	}

	port, err := env.MustGetEnvironmentVariable(pgPortEnvKey)
	if err != nil {
		return postgresConfig{}, errBuildingPostgresConfig(err)
	}

	name, err := env.MustGetEnvironmentVariable(pgNameEnvKey)
	if err != nil {
		return postgresConfig{}, errBuildingPostgresConfig(err)
	}

	username, err := env.MustGetEnvironmentVariable(pgUsernameEnvKey)
	if err != nil {
		return postgresConfig{}, errBuildingPostgresConfig(err)
	}

	password, err := env.MustGetEnvironmentVariable(pgPasswordEnvKey)
	if err != nil {
		return postgresConfig{}, errBuildingPostgresConfig(err)
	}

	return postgresConfig{
		host:     host,
		port:     port,
		name:     name,
		username: username,
		password: password,
	}, nil
}
