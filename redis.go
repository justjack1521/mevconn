package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal"
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

type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

func (c RedisConfig) DSN() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func NewRedisConfig() (RedisConfig, error) {
	host, err := internal.MustGetEnvironmentVariable(redisHostEnvKey)
	if err != nil {
		return RedisConfig{}, errBuildingRedisConfig(err)
	}

	port, err := internal.MustGetEnvironmentVariable(redisHostPortKey)
	if err != nil {
		return RedisConfig{}, errBuildingRedisConfig(err)
	}
	password, err := internal.MustGetEnvironmentVariable(redisHostPasswordKey)
	if err != nil {
		return RedisConfig{}, errBuildingRedisConfig(err)
	}
	return RedisConfig{
		Host:     host,
		Port:     port,
		Password: password,
	}, nil
}
