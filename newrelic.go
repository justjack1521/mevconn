package mevconn

import (
	"fmt"
	"github.com/justjack1521/mevconn/internal/env"
)

const (
	newRelicGuidEnvKey    = "NEWRELICAPPGUID"
	newRelicNameEnvKey    = "NEWRELICAPPNAME"
	newRelicLicenseEnvKey = "NEWRELICLICENSEKEY"
)

var (
	errBuildingNewRelicConfig = func(err error) error {
		return fmt.Errorf("error when building newrelic config: %w", err)
	}
)

type NewRelicConfig interface {
	ApplicationGUID() string
	ApplicationName() string
	LicenseKey() string
}

type newRelicConfig struct {
	appName    string
	appGuid    string
	licenseKey string
}

func (c newRelicConfig) ApplicationName() string {
	return c.appName
}

func (c newRelicConfig) ApplicationGUID() string {
	return c.appGuid
}

func (c newRelicConfig) LicenseKey() string {
	return c.licenseKey
}

func CreateNewRelicConfig() (NewRelicConfig, error) {

	app, err := env.MustGetEnvironmentVariable(newRelicNameEnvKey)
	if err != nil {
		return newRelicConfig{}, errBuildingNewRelicConfig(err)
	}

	key, err := env.MustGetEnvironmentVariable(newRelicLicenseEnvKey)
	if err != nil {
		return newRelicConfig{}, errBuildingNewRelicConfig(err)
	}

	return newRelicConfig{
		appName:    app,
		licenseKey: key,
	}, nil

}
