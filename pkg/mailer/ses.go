package mailer

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/fatkulnurk/gostarter/pkg/interfaces"
)

type SESMailer struct {
	client *ses.Client
	from   string
}

func NewSESMailer(client *ses.Client, from string) interfaces.IMailer {
	return &SESMailer{
		client: client,
		from:   from,
	}
}

func (m *SESMailer) SendMail(msg interfaces.MailMessage) error {
	input := &ses.SendEmailInput{
		Destination: &types.Destination{
			ToAddresses: []string{msg.To},
		},
		Message: &types.Message{
			Body: &types.Body{},
			Subject: &types.Content{
				Data: &msg.Subject,
			},
		},
		Source: &m.from,
	}

	if msg.IsHTML {
		input.Message.Body.Html = &types.Content{Data: &msg.Body}
	} else {
		input.Message.Body.Text = &types.Content{Data: &msg.Body}
	}

	_, err := m.client.SendEmail(context.Background(), input)
	return err
}
