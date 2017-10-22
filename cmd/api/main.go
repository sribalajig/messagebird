package main

import (
	"fmt"
	"log"
	"messagebird/pkg/api"
	"messagebird/pkg/service"
	"os"
)

const (
	envPort   = "HTTP_PORT"
	envAPIKey = "MESSAGE_BIRD_API_KEY"
)

func main() {
	apiKey := os.Getenv(envAPIKey)
	if apiKey == "" {
		log.Println(fmt.Sprintf("%s is not defined in environment variables.", envAPIKey))
		os.Exit(1)
	}

	server := api.NewServer(service.NewMessageBirdAdapter(apiKey))

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
