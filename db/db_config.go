package db

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	dbPath string
}

func newDbConfig() *Config {
	err := godotenv.Load("env_variables.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbConfig := Config{}
	dbConfig.dbPath = os.Getenv("CSV_PATH")

	return &dbConfig
}