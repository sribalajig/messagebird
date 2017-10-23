package service

import (
	"messagebird/pkg/model"

	uuid "github.com/satori/go.uuid"
)

// SMSService is a generic service for sending SMS's
type SMSService struct {
	workerPool *workerPool
}

// NewSMSService returns a pointer to SMSService
func NewSMSService(adapter *MessageBirdAdapter) *SMSService {
	wp := newWorkerPool(adapter)

	go wp.Init()

	return &SMSService{
		workerPool: wp,
	}
}

// Send - sends the given SMS by taking rate limiting into account
func (service *SMSService) Send(sms model.SMS) {
	sms.Reference = uuid.NewV4().String()

	sp := split(sms)

	for _, s := range sp {
		go service.workerPool.Do(s)
	}
}
