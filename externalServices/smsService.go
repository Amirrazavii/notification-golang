package externalServices

import (
	"fmt"
)

type SmsService struct {
}

func (e *SmsService) SendNotify(receiver string, message string) {
	fmt.Printf("sms sent: receiver: %s, message: %s\n", receiver, message)
}

func NewSmsService() *SmsService {
	return &SmsService{}
}
