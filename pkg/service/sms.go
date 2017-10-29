package service

import (
	"messagebird/pkg/datastore"
	"messagebird/pkg/model"

	uuid "github.com/satori/go.uuid"
)

// SMSService is a generic service for sending SMS's
type SMSService struct {
	workerPool   *workerPool
	dataProvider datastore.Provider
}

// NewSMSService returns a pointer to SMSService
func NewSMSService(adapter *MessageBirdAdapter, provider datastore.Provider) *SMSService {
	wp := newWorkerPool(adapter)

	smsService := &SMSService{
		workerPool:   wp,
		dataProvider: provider,
	}

	go wp.Init(func(res response, err chan<- error) {
		err <- smsService.UpdateStatus(res.Reference, res.Status)
	})

	return smsService
}

// Send - sends the given SMS by taking rate limiting into account
func (service *SMSService) Send(sms model.SMS) string {
	sms.Reference = uuid.NewV4().String()

	sp := split(sms)

	for _, s := range sp {
		service.dataProvider.Create(&s)

		go service.workerPool.Do(s)
	}

	return sms.Reference
}

// Get returns SMS's by reference ID
func (service *SMSService) Get(refID string) []model.SMS {
	return service.dataProvider.GetByRefID(refID)
}

// UpdateStatus updates the status of the given reference ID to the given status
func (service *SMSService) UpdateStatus(refID string, status string) error {
	return service.dataProvider.UpdateStatus(refID, status)
}
