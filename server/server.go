package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	Db "server/ip2country/db"
	"time"
)

type Server struct {
	db          Db.Db
	config      *Config
	rateLimiter chan bool
}

func NewServer() *Server {
	server := Server{}
	server.db = Db.NewDb()
	server.config = newConfig()
	server.rateLimiter = make(chan bool, server.config.maxRequests)

	return &server
}

func (server *Server) Run() {
	go server.sendTick()

	http.HandleFunc(server.config.findCountryEndpoint, server.handleGetLocation)

	fmt.Printf("Starting server at port %v\n", server.config.appPort)
	if err := http.ListenAndServe(":" + server.config.appPort, nil); err != nil {
		log.Fatal("failed to start server", err)
	}
}

func (server *Server) handleGetLocation(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != server.config.findCountryEndpoint {
		respondWithError(responseWriter, http.StatusNotFound, "404 not found.")
		return
	}

	if request.Method != http.MethodGet {
		respondWithError(responseWriter, http.StatusMethodNotAllowed, "Method is not supported.")
		return
	}

	ip := request.URL.Query().Get("ip")
	select {
	case <-server.rateLimiter:
		log.Printf("received request for ip: %v\n", ip)
		respondWithJSON(responseWriter, http.StatusOK, server.db.GetLocation(ip))
	default:
		log.Printf("rejected request for ip: %v with %v code\n", ip, http.StatusTooManyRequests)
		respondWithError(responseWriter, http.StatusTooManyRequests, "Too many requests.")
	}
}

func (server *Server) sendTick() {
	rate := time.Tick(time.Second * 3)
	for range rate {
		server.rateLimiter <- true
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(response)
	if err != nil {
		log.Fatal("failed to write response to client", err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}