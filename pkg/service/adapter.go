package service

import (
	"messagebird/pkg/model"
	"time"

	"github.com/messagebird/go-rest-api"
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

// Send - send an SMS, will time out in half a second
func (adapter *MessageBirdAdapter) Send(sms model.SMS, resp chan<- response) error {
	client := messagebird.New(adapter.APIKey)

	task := func() error {
		message, err := client.NewMessage(
			sms.Originator,
			[]string{sms.Recipient},
			sms.Message,
			&messagebird.MessageParams{Reference: sms.Reference})

		if len(message.Recipients.Items) > 0 {
			resp <- response{
				Status:    message.Recipients.Items[0].Status,
				Reference: message.Reference,
			}
		}

		return err
	}

	timeout := newTimeout(time.Millisecond * 500)

	return timeout.Do(task)
}
