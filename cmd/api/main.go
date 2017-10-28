package main

import (
	"fmt"
	"log"
	"messagebird/pkg/api"
	"messagebird/pkg/datastore"
	"messagebird/pkg/service"
	"os"
)

const (
	envDbHost   = "DB_HOST"
	envDbPort   = "DB_PORT"
	envHTTPPort = "HTTP_PORT"
	envAPIKey   = "MESSAGE_BIRD_API_KEY"
)

func main() {
	apiKey := os.Getenv(envAPIKey)
	if apiKey == "" {
		log.Println(fmt.Sprintf("%s is not defined in environment variables.", envAPIKey))
		os.Exit(1)
	}

	dbHost := os.Getenv(envDbHost)
	port := os.Getenv(envDbPort)
	if dbHost == "" || port == "" {
		log.Fatal("Database host/port not set in the Environment. Exiting..")
		os.Exit(0)
	}

	sessFactory, err := datastore.NewSessionFactory(dbHost, port)
	if err != nil {
		log.Fatalf("Unable to connect to mongo db : %s. Exiting., host : %s, port %s", err.Error(), "mongo", "27017")
		os.Exit(1)
	}

	server := api.NewServer(
		service.NewSMSService(
			service.NewMessageBirdAdapter(apiKey),
			datastore.NewMongo(sessFactory),
		),
	)

	port = os.Getenv(envHTTPPort)
	if port == "" {
		log.Fatalf("Environment variable '%s' is not set", envHTTPPort)
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
