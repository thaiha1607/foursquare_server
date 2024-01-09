package env

import (
	"fmt"
	"log"
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

func MustGetEnv(key string) (envVar string) {
	envVar, err := GetEnv(key)
	if err != nil {
		log.Fatal(err)
	}
	return envVar
}

func GetEnvOrDefault(key string, defaultValue string) (envVar string) {
	envVar = os.Getenv(key)
	if envVar == "" {
		log.Printf("environment variable %s not set, using default value", key)
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
