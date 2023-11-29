package env

import (
	"log/slog"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("Could not load .env file")
	}
}
