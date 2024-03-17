package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal/env"
)

const (
	redisHostEnvKey      = "REDISHOST"
	redisHostPortKey     = "REDISPORT"
	redisHostPasswordKey = "REDISPASSWORD"
)

var (
	errBuildingRedisConfig = func(err error) error {
		return fmt.Errorf("error when building redis config: %w", err)
	}
)

type RedisConfig interface {
	DSN() string
	Password() string
}

type redisConfig struct {
	host     string
	port     string
	password string
}

func (c redisConfig) DSN() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}

func (c redisConfig) Password() string {
	return c.password
}

func NewRedisConfig() (RedisConfig, error) {
	host, err := env.MustGetEnvironmentVariable(redisHostEnvKey)
	if err != nil {
		return redisConfig{}, errBuildingRedisConfig(err)
	}

	port, err := env.MustGetEnvironmentVariable(redisHostPortKey)
	if err != nil {
		return redisConfig{}, errBuildingRedisConfig(err)
	}
	password, err := env.MustGetEnvironmentVariable(redisHostPasswordKey)
	if err != nil {
		return redisConfig{}, errBuildingRedisConfig(err)
	}
	return redisConfig{
		host:     host,
		port:     port,
		password: password,
	}, nil
}
