package config

import (
	"fmt"
	"os"
)

const DatabaseURL = "DATABASE_URL"
const JwksUrl = "JWKS_URL"
const TokenAudience = "TOKEN_AUDIENCE"

func GetEnvOrPanic(name string) string {
	value := os.Getenv(name)

	if value == "" {
		panic(fmt.Sprintf("Missing environment variable %s", name))
	}

	return value
}

func GetEnvOrDefault(name string, fallback string) string {
	value := os.Getenv(name)

	if value == "" {
		return fallback
	}

	return value
}
