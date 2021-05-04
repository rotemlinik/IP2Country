package main

import (
	"log"
	"os"
	"server/ip2country/server"
)

func main() {
	f, err := os.OpenFile("app.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("error closing file: %v", err)
		}
	}(f)

	log.SetOutput(f)

	s := server.NewServer()
	s.Run()
}
