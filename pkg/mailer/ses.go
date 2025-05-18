package mailer

import (
	"context"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/fatkulnurk/gostarter/config"
	"github.com/fatkulnurk/gostarter/pkg/interfaces"
	"github.com/fatkulnurk/gostarter/pkg/logging"
)

func NewSESClient(cfg *config.SES) (*ses.Client, error) {
	awscfg, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion("us-west-2"))
	if err != nil {
		logging.Fatalf("unable to load SDK config, %v", err)
		return nil, err
	}

	return ses.NewFromConfig(awscfg), nil
}

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
