package services

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/resend/resend-go/v3"
)

type MailerService struct {
	Client *resend.Client
}

func NewMailerService() (*MailerService, error) {
	apiKey := os.Getenv("RESEND_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("RESEND_API_KEY is not set")
	}

	return &MailerService{
		Client: resend.NewClient(apiKey),
	}, nil
}

func (m *MailerService) SendOTP(to string) error {
	_, err := GenerateOTP(6)
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	emailParams := &resend.SendEmailRequest{
		From:    "Mitter <onboarding@ximon.dev>",
		To:      []string{to},
		Html:    "<strong>hello world</strong>",
		Subject: "OTP Verification",
	}

	sent, err := m.Client.Emails.Send(emailParams)
	if err != nil {
		return err
	}

	fmt.Println(sent.Id)
	return nil
}

func GenerateOTP(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	var otp string
	for i := 0; i < length; i++ {
		digit := bytes[i] % 10
		otp += fmt.Sprintf("%d", digit)
	}

	return otp, nil
}
