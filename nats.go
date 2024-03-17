package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal"
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

type NATSConfig struct {
	Host  string
	Port  string
	Token string
}

func (c NATSConfig) URL() string {
	return fmt.Sprintf("nats://%s:%s", c.Host, c.Port)
}

func NewNATSConfig() (NATSConfig, error) {
	host, err := internal.MustGetEnvironmentVariable(natsHostEnvKey)
	if err != nil {
		return NATSConfig{}, errBuildingNATSConfig(err)
	}

	port, err := internal.MustGetEnvironmentVariable(natsPortEnvKey)
	if err != nil {
		return NATSConfig{}, errBuildingNATSConfig(err)
	}

	token, err := internal.MustGetEnvironmentVariable(natsTokenKey)
	if err != nil {
		return NATSConfig{}, errBuildingNATSConfig(err)
	}

	return NATSConfig{
		Host:  host,
		Port:  port,
		Token: token,
	}, nil
}
