package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	MongoURI string
	DBName string
	MailerHost string
	MailerUsername string
	MailerPassword string
	Secret string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	AppConfig = &Config{
		Port: getEnv("PORT", "8000"),
		MongoURI: getEnv("MONGOURI", "localhost:27018"),
		DBName: getEnv("DB_NAME", "go"),
		MailerHost: getEnv("MAILER_HOST", "smtp.gmail.com"),
		MailerUsername: getEnv("MAILER_USERNAME", ""),
		MailerPassword: getEnv("MAILER_PASSWORD", ""),
		Secret: getEnv("SECRET", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
