package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal/env"
)

const (
	jwtAccessSecretEnvKey  = "JWTACCESSSECRET"
	jwtRefreshSecretEnvKey = "JWTREFRESHSECRET"
)

var (
	errBuildingJWTConfig = func(err error) error {
		return fmt.Errorf("error when building jwt config: %w", err)
	}
)

type JWTConfig interface {
	Secrets() (access string, refresh string)
}

type jwtConfig struct {
	accessSecret  string
	refreshSecret string
}

func (c jwtConfig) Secrets() (access string, refresh string) {
	return c.accessSecret, c.refreshSecret
}

func NewJWTConfig() (JWTConfig, error) {
	access, err := env.MustGetEnvironmentVariable(jwtAccessSecretEnvKey)
	if err != nil {
		return jwtConfig{}, errBuildingJWTConfig(err)
	}

	refresh, err := env.MustGetEnvironmentVariable(jwtRefreshSecretEnvKey)
	if err != nil {
		return jwtConfig{}, errBuildingJWTConfig(err)
	}

	return jwtConfig{
		accessSecret:  access,
		refreshSecret: refresh,
	}, nil

}
