package db

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	dbPath string
	dbUserName string
	dbPassword string
}

func newDbConfig() *Config {
	err := godotenv.Load("env_variables.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbConfig := Config{}
	dbConfig.dbPath = os.Getenv("DB_PATH")
	dbConfig.dbUserName = os.Getenv("DB_USER_NAME")
	dbConfig.dbPassword = os.Getenv("DB_PASSWORD")

	return &dbConfig
}