package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal/env"
)

const (
	rmqHostEnvKey     = "RMQHOST"
	rmqPortEnvKey     = "RMQPORT"
	rmqUserNameEnvKey = "RMQUSERNAME"
	rmqPasswordEnvKey = "RMQPASSWORD"
)

var (
	errBuildingRabbitMQConfig = func(err error) error {
		return fmt.Errorf("error when building rabbit mq config: %w", err)
	}
)

type RabbitMQConfig interface {
	Source() string
}

type rabbitMQConfig struct {
	host     string
	port     string
	username string
	password string
}

func (c rabbitMQConfig) Source() string {
	return fmt.Sprintf("amqps://%s:%s@%s/%s", c.username, c.password, c.host, c.username)
}

func CreateRabbitMQConfig() (RabbitMQConfig, error) {

	host, err := env.MustGetEnvironmentVariable(rmqHostEnvKey)
	if err != nil {
		return rabbitMQConfig{}, errBuildingRabbitMQConfig(err)
	}

	port, err := env.MustGetEnvironmentVariable(rmqPortEnvKey)
	if err != nil {
		return rabbitMQConfig{}, errBuildingRabbitMQConfig(err)
	}

	username, err := env.MustGetEnvironmentVariable(rmqUserNameEnvKey)
	if err != nil {
		return rabbitMQConfig{}, errBuildingRabbitMQConfig(err)
	}

	password, err := env.MustGetEnvironmentVariable(rmqPasswordEnvKey)
	if err != nil {
		return rabbitMQConfig{}, errBuildingRabbitMQConfig(err)
	}

	return rabbitMQConfig{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}, nil

}
