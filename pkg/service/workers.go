package service

import (
	"messagebird/pkg/model"
	"time"
)

type workerPool struct {
	sms       chan model.SMS
	rateLimit <-chan time.Time
	done      chan bool
	adapter   *MessageBirdAdapter
}

func newWorkerPool() *workerPool {
	return &workerPool{
		sms:       make(chan (model.SMS)),
		rateLimit: time.Tick(time.Second * 1),
		done:      make(chan (bool)),
		adapter:   &MessageBirdAdapter{},
	}
}

// Init gets the worker pool going
func (w *workerPool) Init() {
	for {
		sms := <-w.sms

		select {
		case <-w.rateLimit:
			go w.adapter.Send(sms)
		case <-w.done:
			break
		}
	}
}

func (w *workerPool) Do(sms model.SMS) {
	w.sms <- sms
}

func (w *workerPool) Shutdown() {
	w.done <- true
}
