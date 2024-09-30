package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal/env"
)

const (
	keyCloakClientIDKey     = "KCCLIENTID"
	keyCloakClientSecretKey = "KCCLIENTSECRET"
	keyCloakRealmKey        = "KCREALM"
)

var (
	errBuildingKeyCloakConfig = func(err error) error {
		return fmt.Errorf("error when building keycloak config: %w", err)
	}
)

type KeyCloakConfig interface {
	ClientID() string
	ClientSecret() string
	Realm() string
}

type keyCloakConfig struct {
	clientID     string
	clientSecret string
	realm        string
}

func (c keyCloakConfig) ClientID() string {
	return c.clientID
}

func (c keyCloakConfig) ClientSecret() string {
	return c.clientSecret
}

func (c keyCloakConfig) Realm() string {
	return c.realm
}

func NewKeyCloakConfig() (KeyCloakConfig, error) {

	id, err := env.MustGetEnvironmentVariable(keyCloakClientIDKey)
	if err != nil {
		return nil, errBuildingKeyCloakConfig(err)
	}

	secret, err := env.MustGetEnvironmentVariable(keyCloakClientSecretKey)
	if err != nil {
		return nil, errBuildingKeyCloakConfig(err)
	}

	realm, err := env.MustGetEnvironmentVariable(keyCloakRealmKey)
	if err != nil {
		return nil, errBuildingKeyCloakConfig(err)
	}

	return keyCloakConfig{
		clientID:     id,
		clientSecret: secret,
		realm:        realm,
	}, nil

}
