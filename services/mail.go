package services

import (
	"fmt"
	"os"

	"github.com/resend/resend-go/v3"
)

type MailService struct {
	client *resend.Client
}

func NewMailService() (*MailService, error) {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("RESEND_API_KEY is not set")
	}

	return &MailService{
		client: resend.NewClient(apiKey),
	}, nil
}

func (m *MailService) SendOTP(to string, otp string) error {
	emailParams := &resend.SendEmailRequest{
		From:    "Mitter <onboarding@ximon.dev>",
		To:      []string{to},
		Html:    "<strong>hello world</strong>",
		Subject: "OTP Verification",
	}

	sent, err := m.client.Emails.Send(emailParams)
	if err != nil {
		return err
	}

	fmt.Println(sent.Id)
	return nil
}
