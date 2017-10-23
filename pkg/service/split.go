package service

import (
	"fmt"
	"messagebird/pkg/model"

	uuid "github.com/satori/go.uuid"
)

func split(sms model.SMS) []model.SMS {
	if len(sms.Message) <= 160 {
		return []model.SMS{sms}
	}

	chunks := (len(sms.Message) / 161) + 1

	if chunks > 9 {
		chunks = 9
	}

	var s []model.SMS
	for i := 0; i < chunks; i++ {
		chunkStart := getChunkStart(i)

		var chunkEnd int
		if i == chunks-1 {
			chunkEnd = len(sms.Message)
		} else {
			chunkEnd = getChunkEnd(i)
		}

		smsChunk := model.SMS{
			Reference:  uuid.NewV4().String(),
			Recipient:  sms.Recipient,
			Originator: sms.Originator,
			Message:    sms.Message[chunkStart:chunkEnd],
			UDH:        fmt.Sprintf("050003A6%s%s", digits[chunks], digits[i]),
		}

		s = append(s, smsChunk)
	}

	return s
}

func getChunkStart(index int) int {
	if index == 0 {
		return 0
	}

	return 160*index + 1
}

func getChunkEnd(index int) int {
	return (index+1)*160 + 1
}

var digits = map[int]string{
	1: "01",
	2: "02",
	3: "03",
	4: "04",
	5: "05",
	6: "06",
	7: "07",
	8: "08",
	9: "09",
}
