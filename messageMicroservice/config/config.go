package config

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
)

var l = hclog.Default().Named("Config")

func verifyAllEnvVars() {
	envVars := []string{"SENDGRID_API_KEY"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			l.Error(("Missing env var: " + envVar))
		}
	}
}

func init() {
	godotenv.Load(".env")
	verifyAllEnvVars()
}

type Config struct {
	SENDGRID_API_KEY string
	RABBIT_URL       string
}

func New() *Config {
	return &Config{
		SENDGRID_API_KEY: os.Getenv("SENDGRID_API_KEY"),
		RABBIT_URL:       os.Getenv("RABBIT_URL"),
	}
}
