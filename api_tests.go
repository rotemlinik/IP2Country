package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load("env_variables.env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }

	findCountryEndpoint = os.Getenv("API_FIND_COUNTRY_ENDPOINT")
	appPort = os.Getenv("APP_PORT")
	maxRequests, err := strconv.Atoi(os.Getenv("MAX_REQUESTS_PER_SECOND"))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	rateLimiter = make(chan bool, maxRequests)
}

func TestSomething(t *testing.T) {
  // assert equality
  assert.Equal(t, 123, 123, "they should be equal")

  // assert inequality
  assert.NotEqual(t, 123, 456, "they should not be equal")

  // assert for nil (good for errors)
  assert.Nil(t, object)

  // assert for not nil (good when you expect something)
  if assert.NotNil(t, object) {

    // now we know that object isn't nil, we are safe to make
    // further assertions without causing any errors
    assert.Equal(t, "Something", object.Value)

  }
}