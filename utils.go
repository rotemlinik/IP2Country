package main

import (
	"log"
)

func HandleFailure(msg string, err error) {
	log.Fatalf(msg, err) //TODO where is this written to? add log file if needed
}