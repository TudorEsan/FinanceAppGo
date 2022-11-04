package config

import (
	"os"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/joho/godotenv"
)

var l = hclog.Default().Named("Config")

func verifyAllEnvVars() {
	envVars := []string{"MONGO_URL", "JWT_SECRET", "SENDGRID_API_KEY", "DOMAIN_NAME"}
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
	MongoUrl      string
	EncryptionKey []byte
	MongoTimeout  time.Duration
	RabbitUrl     string
	RedisUrl      string
	RedisPassword string
}

func getConfig() *Config {
	// defaults
	mongoUrl := "mongodb://localhost:27017"
	mongoTimeout := time.Second * 10

	if os.Getenv("TESTING") != "" {
		goto ReturnConfig
	}

	if os.Getenv("MONGO_URL") != "" {
		mongoUrl = os.Getenv("MONGO_URL")
	}

	if os.Getenv("ENCRYPTION_KEY") == "" {
		panic("Encryption key is not set")
	}

	if os.Getenv("REDIS_URL") == "" {
		panic("Redis URL is not set")
	}

ReturnConfig:
	return &Config{
		MongoUrl:      mongoUrl,
		EncryptionKey: []byte(os.Getenv("ENCRYPTION_KEY")),
		MongoTimeout:  mongoTimeout,
		RabbitUrl:     os.Getenv("RABBIT_URL"),
		RedisUrl:      os.Getenv("REDIS_URL"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
	}
}

func New() *Config {
	return getConfig()
}
