package config

import (
	"os"

	"github.com/joho/godotenv"
)

func verifyAllEnvVars() {
	envVars := []string{"MONGO_URL", "JWT_SECRET", "SENDGRID_API_KEY", "DOMAIN_NAME"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			panic("Missing env var: " + envVar)
		}
	}
}

func init() {
	godotenv.Load(".env")
	verifyAllEnvVars()
}

type Config struct {
	MongoUrl     string
	JwtSecret    []byte
	SmtpUsername string
	SmtpPassword string
}

func getConfig() *Config {
	return &Config{
		MongoUrl:     os.Getenv("MONGO_URL"),
		JwtSecret:    []byte(os.Getenv("JWT_SECRET")),
		SmtpUsername: os.Getenv("SMTP_USERNAME"),
		SmtpPassword: os.Getenv("SMTP_PASSWORD"),
	}
}

func New() *Config {
	return getConfig()
}
