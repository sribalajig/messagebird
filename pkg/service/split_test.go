package service

import (
	"fmt"
	"messagebird/pkg/model"
	"testing"
)

func TestSplitTwoSMSs(t *testing.T) {
	// msg has a length of 174
	msg := `I really like showing what shocks people. 
			I know that people don't like to tell the 
			truth all the time, but what people don't 
			like to say is exactly what drives me to paint.`

	sms := model.SMS{
		Message: msg,
	}

	sp := split(sms)

	if len(sp) != 2 {
		t.Fatalf("Expected %d messages, got %d", 2, len(sp))
	}

	expected := fmt.Sprintf("%s%s", sp[0].Message, sp[1].Message)
	if expected != msg {
		t.Fatalf("The split messages don't add up to the original original '%s', split '%s'",
			msg,
			expected)
	}
}

func TestSplitOneSMS(t *testing.T) {
	// msg has a length of 160
	msg := `We've forgotten much. How to struggle,
			how to rise to dizzy heights and sink
			to unparalleled depths. We no longer
			aspire to anything. Even the finer`

	sms := model.SMS{
		Message: msg,
	}

	sp := split(sms)

	if len(sp) != 1 {
		t.Fatalf("Expected %d messages, got %d", 1, len(sp))
	}
}

func TestSplitThreeSMSs(t *testing.T) {
	msg := `He had come to that moment in his age when 
	there occurred to him, with increasing intensity, 
	a question of such overwhelming simplicity that he 
	had no means to face it. He found himself wondering 
	if his life were worth the living; if it had ever been. 
	It was a question, he suspected, that came to all men 
	at one time or another;`

	sms := model.SMS{
		Message: msg,
	}

	sp := split(sms)

	expected := fmt.Sprintf("%s%s%s", sp[0].Message, sp[1].Message, sp[2].Message)
	if expected != msg {
		t.Fatalf("The split messages don't add up to the original original '%s', split '%s'",
			msg,
			fmt.Sprintf("%s%s%s", sp[0].Message, sp[1].Message, sp[2].Message))
	}

	if sp[0].UDH != "050003A60301" {
		t.Fatalf("Expected '%s', got '%s'", "050003A60301", sp[0].UDH)
	}

	if sp[1].UDH != "050003A60302" {
		t.Fatalf("Expected '%s', got '%s'", "050003A60302", sp[0].UDH)
	}

	if sp[2].UDH != "050003A60303" {
		t.Fatalf("Expected '%s', got '%s'", "050003A60303", sp[0].UDH)
	}
}
