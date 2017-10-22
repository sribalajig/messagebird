package service

import (
	"errors"
	"fmt"
	"log"
	"messagebird/pkg/model"

	"time"

	"github.com/messagebird/go-rest-api"
	"github.com/satori/go.uuid"
)

// MessageBirdAdapter sends messages using the messagebird API
type MessageBirdAdapter struct {
	APIKey string
}

// NewMessageBirdAdapter returns a pointer to MessageBirdAdapter
func NewMessageBirdAdapter(apiKey string) *MessageBirdAdapter {
	return &MessageBirdAdapter{
		APIKey: apiKey,
	}
}

// Send ...
func (adapter *MessageBirdAdapter) Send(sms model.SMS) error {
	client := messagebird.New(adapter.APIKey)

	ref := uuid.NewV4()

	timeout := time.Tick(time.Millisecond * 500)
	done := make(chan (bool))
	var err error
	go func(chan<- bool) {
		message, err := client.NewMessage(
			sms.Originator,
			[]string{sms.Recipient},
			sms.Message,
			&messagebird.MessageParams{Reference: ref.String()})

		log.Println(fmt.Sprintf("This is the response from the messagebird api : %#v", message))

		if err != nil {
			log.Println(fmt.Sprintf("Error while sending SMS through messagebird api : '%s'", err.Error()))
		}

		done <- true
	}(done)

	select {
	case <-done:
		return err
	case <-timeout:
		return errors.New("HTTP request to message bir api timed out")
	}
}
