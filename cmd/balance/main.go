package main

import (
	"fmt"
	"log"
	"os"

	"github.com/messagebird/go-rest-api"
)

const envAPIKey = "MESSAGE_BIRD_API_KEY"

func main() {
	apiKey := os.Getenv(envAPIKey)
	if apiKey == "" {
		log.Println(fmt.Sprintf("%s is not defined in environment variables.", envAPIKey))
		os.Exit(1)
	}

	client := messagebird.New(apiKey)

	balance, err := client.Balance()
	if err != nil {
		// messagebird.ErrResponse means custom JSON errors.
		if err == messagebird.ErrResponse {
			for _, mbError := range balance.Errors {
				fmt.Printf("Error: %#v\n", mbError)
			}
		}

		return
	}

	fmt.Println("  payment :", balance.Payment)
	fmt.Println("  type    :", balance.Type)
	fmt.Println("  amount  :", balance.Amount)
}
