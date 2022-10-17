package config

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
)

var l = hclog.Default().Named("Config")

func verifyAllEnvVars() {
	envVars := []string{"SENDGRID_API_KEY", "RABBIT_URL"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			l.Error(("Missing env var: " + envVar))
		}
	}
}

var loaded = false

func init() {
	loadEnvs()
}

func loadEnvs() {
	l.Info("Loading env vars")
	err := godotenv.Load("../.env")
	if err != nil {
		l.Error("Error loading .env file")
	}
	loaded = true
	verifyAllEnvVars()
}

type Config struct {
	SENDGRID_API_KEY string
	RABBIT_URL       string
}

func New() *Config {
	if !loaded {
		loadEnvs()
	}
	return &Config{
		SENDGRID_API_KEY: os.Getenv("SENDGRID_API_KEY"),
		RABBIT_URL:       os.Getenv("RABBIT_URL"),
	}
}
