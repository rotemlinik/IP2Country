package main

import (
  "fmt"
  "log"
  "net/http"
	"time"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var rateLimiter chan bool
var findCountryEndpoint string
var appPort string

func sendTick(rateLimiter chan<- bool) {
	rate := time.Tick(time.Second)
	for range rate {
		rateLimiter <- true
	}
}

func findCountryHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != findCountryEndpoint {
			http.Error(responseWriter, "404 not found.", http.StatusNotFound)
			return
	}

	if request.Method != http.MethodGet {
			http.Error(responseWriter, "Method is not supported.", http.StatusMethodNotAllowed)
			return
	}

	select {
		case <-rateLimiter:
			fmt.Println("processing request!")
			ip := request.URL.Query().Get("ip")
			fmt.Println("ip =>", ip)
		default:
			http.Error(responseWriter, "Too many requests.", http.StatusTooManyRequests)
	}
}

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

func main() {
	go sendTick(rateLimiter)

  http.HandleFunc(findCountryEndpoint, findCountryHandler)

	fmt.Printf("Starting server at port %v\n", appPort)
	if err := http.ListenAndServe(":" + appPort, nil); err != nil {
		//TODO maybe exit here aswell?
		log.Fatal(err)
	}
}