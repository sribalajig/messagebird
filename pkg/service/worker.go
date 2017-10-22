package service

import (
	"fmt"
	"log"
	"messagebird/pkg/model"
	"time"
)

type workerPool struct {
	sms       chan model.SMS
	rateLimit <-chan time.Time
	resp      chan response
	done      chan bool
	adapter   *MessageBirdAdapter
}

func newWorkerPool(adapter *MessageBirdAdapter) *workerPool {
	return &workerPool{
		sms:       make(chan (model.SMS)),
		rateLimit: time.Tick(time.Second * 1),
		done:      make(chan (bool)),
		resp:      make(chan (response)),
		adapter:   adapter,
	}
}

// Init gets the worker pool going
func (w *workerPool) Init() {
	for {
		select {
		case <-w.rateLimit:
			sms := <-w.sms
			log.Println("Sending the request now!")
			go w.adapter.Send(sms, w.resp)
		case rs := <-w.resp:
			log.Println("Have to now process the response.")
			log.Println(fmt.Sprintf("Message : %#v", rs))
		case <-w.done:
			break
		}
	}
}

func (w *workerPool) Do(sms model.SMS) {
	log.Println("Pushing the message..")
	w.sms <- sms
}

func (w *workerPool) Shutdown() {
	w.done <- true
}
