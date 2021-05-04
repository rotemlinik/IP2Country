package server

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	findCountryEndpoint string
	appPort             string
	maxRequests         int
}

func newConfig() *Config {
	apiConfig := Config{}
	err := godotenv.Load("env_variables.env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	apiConfig.findCountryEndpoint = os.Getenv("API_FIND_COUNTRY_ENDPOINT")
	apiConfig.appPort = os.Getenv("APP_PORT")
	apiConfig.maxRequests, err = strconv.Atoi(os.Getenv("MAX_REQUESTS_PER_SECOND"))
	if err != nil {
		log.Fatal("Failed to parse environment variable", err)
	}

	return &apiConfig
}
