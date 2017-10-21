package main

import (
	"log"
	"messagebird/pkg/api"
	"os"
)

const envPort = "HTTP_PORT"

func main() {
	server := api.NewServer()

	port := os.Getenv(envPort)
	if port == "" {
		log.Fatalf("Environment variable '%s' is not set", envPort)
		os.Exit(1)
	}

	ch := make(chan error)
	server.Start(port, ch)

	if err := <-ch; err != nil {
		close(ch)
		log.Fatalf("Shutting down HTTP server, there was an error : '%s'", err.Error())
		os.Exit(1)
	} else if err == nil {
		close(ch)
		log.Println("Exiting. No error.")
		os.Exit(0)
	}
}
