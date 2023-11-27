package env

import (
	"fmt"
	"log/slog"
	"os"
)

func GetEnv(key string) string {
	env_var := os.Getenv(key)
	if env_var == "" {
		panic(fmt.Sprintf("Environment variable %s not set", key))
	}
	return env_var
}

func GetEnvOrDefault(key string, default_value string) string {
	env_var := os.Getenv(key)
	if env_var == "" {
		slog.Warn(fmt.Sprintf("Environment variable %s not set, using default value %s", key, default_value))
		return default_value
	}
	return env_var
}

func TryGetEnv(key string) bool {
	env_var := os.Getenv(key)
	return env_var != ""
}
