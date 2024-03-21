package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal/env"
)

const (
	redisHostEnvKey      = "REDISHOST"
	redisHostPortKey     = "REDISPORT"
	redisUserNameKey     = "REDISUSERNAME"
	redisHostPasswordKey = "REDISPASSWORD"
)

var (
	errBuildingRedisConfig = func(err error) error {
		return fmt.Errorf("error when building redis config: %w", err)
	}
)

type RedisConfig interface {
	Host() string
	DSN() string
	Username() string
	Password() string
}

type redisConfig struct {
	host     string
	port     string
	username string
	password string
}

func (c redisConfig) Host() string {
	return c.host
}

func (c redisConfig) Username() string {
	return c.username
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
	user, err := env.MustGetEnvironmentVariable(redisUserNameKey)
	if err != nil {
		return redisConfig{}, err
	}
	password, err := env.MustGetEnvironmentVariable(redisHostPasswordKey)
	if err != nil {
		return redisConfig{}, errBuildingRedisConfig(err)
	}
	return redisConfig{
		host:     host,
		port:     port,
		username: user,
		password: password,
	}, nil
}
