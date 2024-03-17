package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal"
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

type JWTConfig struct {
	AccessSecret  string
	RefreshSecret string
}

func NewJWTConfig() (JWTConfig, error) {
	access, err := internal.MustGetEnvironmentVariable(jwtAccessSecretEnvKey)
	if err != nil {
		return JWTConfig{}, errBuildingJWTConfig(err)
	}

	refresh, err := internal.MustGetEnvironmentVariable(jwtRefreshSecretEnvKey)
	if err != nil {
		return JWTConfig{}, errBuildingJWTConfig(err)
	}

	return JWTConfig{
		AccessSecret:  access,
		RefreshSecret: refresh,
	}, nil

}
