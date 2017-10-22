package service

import (
	"messagebird/pkg/model"
)

// SMSService is a generic service for sending SMS's
type SMSService struct {
	workerPool *workerPool
}

// NewSMSService returns a pointer to SMSService
func NewSMSService() *SMSService {
	wp := newWorkerPool()

	go wp.Init()

	return &SMSService{
		workerPool: wp,
	}
}

// Send - sends the given SMS by taking rate limiting into account
func (service *SMSService) Send(sms model.SMS) {
	go service.workerPool.Do(sms)
}
