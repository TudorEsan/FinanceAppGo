package config

import "os"

type Config struct {
	MongoUrl string
	JwtSecret string
	SmtpUsername string
	SmtpPassword string
}

func getConfig() Config {
	return Config{
		MongoUrl: os.Getenv("MONGO_URL"),
		JwtSecret: os.Getenv("JWT_SECRET"),
		SmtpUsername: os.Getenv("SMTP_USERNAME"),
		SmtpPassword: os.Getenv("SMTP_PASSWORD"),
	}
}

func New() Config {
	return getConfig()
}