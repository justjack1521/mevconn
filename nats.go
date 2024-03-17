package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal/env"
)

const (
	natsHostEnvKey = "NATSHOST"
	natsPortEnvKey = "NATSPORT"
	natsTokenKey   = "NATSTOKEN"
)

var (
	errBuildingNATSConfig = func(err error) error {
		return fmt.Errorf("error when building nats config: %w", err)
	}
)

type NATSConfig interface {
	URL() string
	Token() string
}

type natsConfig struct {
	host  string
	port  string
	token string
}

func (c natsConfig) Token() string {
	return c.token
}

func (c natsConfig) URL() string {
	return fmt.Sprintf("nats://%s:%s", c.host, c.port)
}

func NewNATSConfig() (NATSConfig, error) {
	host, err := env.MustGetEnvironmentVariable(natsHostEnvKey)
	if err != nil {
		return natsConfig{}, errBuildingNATSConfig(err)
	}

	port, err := env.MustGetEnvironmentVariable(natsPortEnvKey)
	if err != nil {
		return natsConfig{}, errBuildingNATSConfig(err)
	}

	token, err := env.MustGetEnvironmentVariable(natsTokenKey)
	if err != nil {
		return natsConfig{}, errBuildingNATSConfig(err)
	}

	return natsConfig{
		host:  host,
		port:  port,
		token: token,
	}, nil
}
