package env

import (
	"fmt"
	"os"
)

var (
	errEnvironmentVariableNotFound = func(key string) error {
		return fmt.Errorf("environment variable not found for key: %s", key)
	}
)

func GetEnvironmentVariable(key string) string {
	return os.Getenv(key)
}

func MustGetEnvironmentVariable(key string) (string, error) {
	var result = GetEnvironmentVariable(key)
	if result == "" {
		return "", errEnvironmentVariableNotFound(key)
	}
	return result, nil
}
