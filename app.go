package main

import "server/ip2country/server"

func main() {
	s := server.NewServer()
	s.Run()
}
