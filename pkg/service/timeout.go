package service

import (
	"fmt"
	"time"
)

type timeout struct {
	duration time.Duration
}

func newTimeout(d time.Duration) *timeout {
	return &timeout{
		duration: d,
	}
}

func (t *timeout) Do(f func() error) error {
	timeout := time.Tick(t.duration)
	done := make(chan (bool))
	var err error

	go func() {
		err = f()
		done <- true
	}()

	select {
	case <-done:
		return err
	case <-timeout:
		return fmt.Errorf("Timeout after %#v", t.duration)
	}
}
