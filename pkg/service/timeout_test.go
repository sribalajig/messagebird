package service

import (
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	timeout := newTimeout(time.Millisecond * 1)

	err := timeout.Do(func() error {
		time.Sleep(time.Millisecond * 5)
		return nil
	})
	if err == nil {
		t.Fatal("Expected error, got nil")
	} else if err.Error() == "" {
		t.Fatalf("Expected an error message, but got empty string")
	}

	timeout = newTimeout(time.Millisecond * 2)
	err = timeout.Do(func() error {
		time.Sleep(time.Millisecond * 1)
		return nil
	})
	if err != nil {
		t.Fatalf("Expected no error, got '%s'", err.Error())
	}
}
