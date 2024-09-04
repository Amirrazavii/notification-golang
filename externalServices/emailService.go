package externalServices

import (
	"fmt"
)

type EmailService struct {
}


func (e *EmailService) SendNotify(receiver string, message string) {
	fmt.Printf("Email sent: receiver: %s, message: %s\n", receiver, message)
}

func NewEmailService() *EmailService {
	return &EmailService{}
}
