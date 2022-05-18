package config

import "os"

type Config struct {
	SERVER_ADDRESS        string
	DB_USERNAME           string
	DB_PASSWORD           string
	DB_PORT               string
	DB_HOST               string
	DB_NAME               string
	MONGO_USERNAME        string
	MONGO_PASSWORD        string
	MONGO_HOST            string
	MONGO_PORT            string
	MONGO_DB_NAME         string
	REDIS_PASSWORD        string
	REDIS_HOST            string
	REDIS_PORT            string
	// REDIS_DB       string
	JWT_SECRET     string
	REFRESH_SECRET string
}

func InitializeConfig() Config {
	return Config{
		SERVER_ADDRESS:        os.Getenv("SERVER_ADDRESS"),
		DB_USERNAME:           os.Getenv("DB_USERNAME"),
		DB_PASSWORD:           os.Getenv("DB_PASSWoRD"),
		DB_PORT:               os.Getenv("DB_PORT"),
		DB_HOST:               os.Getenv("DB_HOST"),
		DB_NAME:               os.Getenv("DB_NAME"),
		MONGO_USERNAME:        os.Getenv("MONGO_USERNAME"),
		MONGO_PASSWORD:        os.Getenv("MONGO_PASSWORD"),
		MONGO_HOST:            os.Getenv("MONGO_HOST"),
		MONGO_PORT:            os.Getenv("MONGO_PORT"),
		MONGO_DB_NAME:         os.Getenv("MONGO_DB_NAME"),
		REDIS_PASSWORD:        os.Getenv("REDIS_PASSWORD"),
		REDIS_HOST:            os.Getenv("REDIS_HOST"),
		REDIS_PORT:            os.Getenv("REDIS_PORT"),
		// REDIS_DB:       os.Getenv("REDIS_DB"),
		JWT_SECRET:     os.Getenv("JWT_SECRET"),
		REFRESH_SECRET: os.Getenv("REFRESH_SECRET"),
	}
}
