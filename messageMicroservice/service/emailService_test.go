package service

import (
	"testing"

	"github.com/TudorEsan/FinanceAppGo/messageMicroservice/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	conf := config.New()
	client := sendgrid.NewSendClient(conf.SENDGRID_API_KEY)
	emailService := NewEmailService(client)
	err := emailService.sendEmail(EmailOptions{
		To:          "Tudor",
		Email:       "tudor.esan@icloud.com",
		Subject:     "Test",
		HtmlContent: "Test",
		Content:     "Test",
	})

	assert.Nil(t, err)

}
