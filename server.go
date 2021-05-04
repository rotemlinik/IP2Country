package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	Db "server/ip2country/db"
	"time"
)

type Server struct {
	db Db.Db
	apiConfig *ApiConfig
	rateLimiter chan bool
}

func newServer() *Server {
	server := Server{}
	server.db = Db.NewDb() // TODO also pointer???
	server.apiConfig = newApiConfig()
	server.rateLimiter = make(chan bool, server.apiConfig.maxRequests)

	return &server
}

func (server *Server) run() {
	go sendTick(server.rateLimiter)

	http.HandleFunc(server.apiConfig.findCountryEndpoint, server.handleGetLocation)

	fmt.Printf("Starting server at port %v\n", server.apiConfig.appPort)
	if err := http.ListenAndServe(":" + server.apiConfig.appPort, nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func sendTick(rateLimiter chan<- bool) {
	rate := time.Tick(time.Second)
	for range rate {
		rateLimiter <- true
	}
}

func (server *Server) handleGetLocation(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != server.apiConfig.findCountryEndpoint {
		respondWithError(responseWriter, http.StatusNotFound, "404 not found.")
	}

	if request.Method != http.MethodGet {
		respondWithError(responseWriter, http.StatusMethodNotAllowed, "Method is not supported.")
	}

	select {
	case <-server.rateLimiter:
		fmt.Println("processing request!")
		ip := request.URL.Query().Get("ip")
		fmt.Println("ip =>", ip)
		respondWithJSON(responseWriter, http.StatusOK, server.db.GetLocation(ip))
	default:
		respondWithError(responseWriter, http.StatusTooManyRequests, "Too many requests.")
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}