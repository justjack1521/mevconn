package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal"
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

type PostgresConfig struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

func (c PostgresConfig) Source() string {
	return fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=require", c.Port, c.Host, c.Username, c.Password, c.Name)
}

func NewPostgresConfig() (PostgresConfig, error) {

	host, err := internal.MustGetEnvironmentVariable(pgHostEnvKey)
	if err != nil {
		return PostgresConfig{}, errBuildingPostgresConfig(err)
	}

	port, err := internal.MustGetEnvironmentVariable(pgPortEnvKey)
	if err != nil {
		return PostgresConfig{}, errBuildingPostgresConfig(err)
	}

	name, err := internal.MustGetEnvironmentVariable(pgNameEnvKey)
	if err != nil {
		return PostgresConfig{}, errBuildingPostgresConfig(err)
	}

	username, err := internal.MustGetEnvironmentVariable(pgUsernameEnvKey)
	if err != nil {
		return PostgresConfig{}, errBuildingPostgresConfig(err)
	}

	password, err := internal.MustGetEnvironmentVariable(pgPasswordEnvKey)
	if err != nil {
		return PostgresConfig{}, errBuildingPostgresConfig(err)
	}

	return PostgresConfig{
		Host:     host,
		Port:     port,
		Name:     name,
		Username: username,
		Password: password,
	}, nil
}
