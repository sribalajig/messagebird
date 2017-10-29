package model

import (
	"errors"
	"strings"
)

// SMS is the data structure for an SMS
type SMS struct {
	Reference  string
	Recipient  string
	Originator string
	Message    string
	UDH        string
	Status     string
}

const (
	// StatusQueued - message is queued to be sent to the messagebird API
	StatusQueued = "queued"
	// StatusSent - message was sent successfully with the messagebird API
	StatusSent = "sent"
)

// IsValid validates the SMS and returns error if any
func (sms *SMS) IsValid() (bool, error) {
	var errs []string

	if sms.Recipient == "" {
		errs = append(errs, "recipient is empty")
	}

	if sms.Originator == "" {
		errs = append(errs, "originator is empty")
	}

	if sms.Message == "" {
		errs = append(errs, "message is empty")
	}

	if len(errs) > 0 {
		err := strings.Join(errs, ",")
		return false, errors.New(err)
	}

	return true, nil
}
