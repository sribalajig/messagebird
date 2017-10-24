package model

import "testing"

var testCases = []struct {
	SMS           SMS
	IsValid       bool
	ExpectedError string
}{
	{SMS{Recipient: "", Originator: "", Message: ""}, false, "recipient is empty,originator is empty,message is empty"},
}

func TestIsValid(t *testing.T) {
	for _, testCase := range testCases {
		isValid, err := testCase.SMS.IsValid()

		if testCase.ExpectedError == "" && err != nil {
			t.Logf("Expected no error but got error '%s'", err.Error())
			t.Fail()
		}

		if testCase.ExpectedError != "" {
			if err == nil {
				t.Logf("Expected error '%s' but got no error", testCase.ExpectedError)
				t.Fail()
			} else if testCase.ExpectedError != err.Error() {
				t.Fail()
				t.Logf("Expected error '%s', got error '%s'", testCase.ExpectedError, err.Error())
			}
		}

		if isValid != testCase.IsValid {
			t.Logf("Expected %t, got %t", testCase.IsValid, isValid)
			t.Fail()
		}
	}
}
