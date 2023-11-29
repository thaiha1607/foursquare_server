package env

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

func GetEnv(key string) (envVar string, err error) {
	envVar = os.Getenv(key)
	if envVar == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return envVar, nil
}

func GetEnvOrDefault(key string, defaultValue string) (envVar string) {
	envVar = os.Getenv(key)
	if envVar == "" {
		slog.Warn(fmt.Sprintf("Environment variable %s not set, using default value %s", key, defaultValue))
		return defaultValue
	}
	return envVar
}

func TryGetEnv(key string) bool {
	return os.Getenv(key) != ""
}

func IsProdEnv() bool {
	return strings.ToLower(GetEnvOrDefault("GO_ECHO_ENV", "")) == "production"
}
