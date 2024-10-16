package utils

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config map[string]*string

func ParseConfig(requestedConfig map[string]string) Config {
	if err := godotenv.Load(); err != nil {

		slog.Error("Error loading .env file")
		os.Exit(1)
	}
	var config Config = make(map[string]*string)

	for name, defaultValue := range requestedConfig {
		var envVar string
		if defaultValue != "" {
			envVar = getEnv(name, &defaultValue)
		}
		envVar = getEnv(name, nil)

		config[name] = &envVar
	}

	return config
}

func getEnv(varName string, defaultValue *string) string {
	envVar := os.Getenv(varName)
	if envVar != "" {
		return envVar
	}

	if defaultValue != nil {
		return *defaultValue
	}

	slog.Error("Missing Env Var" + varName)
	os.Exit(1)
	return ""
}
