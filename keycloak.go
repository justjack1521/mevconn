package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal/env"
)

const (
	keyCloakHostNameKey     = "KCHOSTNAME"
	keyCloakClientIDKey     = "KCCLIENTID"
	keyCloakClientSecretKey = "KCCLIENTSECRET"
	keyCloakRealmKey        = "KCREALM"
)

const (
	keyCloakAdminUsernameKey = "KCADMINUSERNAME"
	keyCloakAdminPasswordKey = "KCADMINPASSWORD"
)

var (
	errBuildingKeyCloakConfig = func(err error) error {
		return fmt.Errorf("error when building keycloak config: %w", err)
	}
)

type KeyCloakConfig interface {
	Hostname() string
	ClientID() string
	ClientSecret() string
	Realm() string
	AdminCredentials() (username string, password string)
}

type keyCloakConfig struct {
	hostName      string
	clientID      string
	clientSecret  string
	realm         string
	adminUsername string
	adminPassword string
}

func (c keyCloakConfig) AdminCredentials() (username string, password string) {
	return c.adminUsername, c.adminPassword
}

func (c keyCloakConfig) Hostname() string {
	return c.hostName
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

	host, err := env.MustGetEnvironmentVariable(keyCloakHostNameKey)
	if err != nil {
		return nil, errBuildingKeyCloakConfig(err)
	}

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
		hostName:      host,
		clientID:      id,
		clientSecret:  secret,
		realm:         realm,
		adminUsername: env.GetEnvironmentVariable(keyCloakAdminUsernameKey),
		adminPassword: env.GetEnvironmentVariable(keyCloakAdminPasswordKey),
	}, nil

}
