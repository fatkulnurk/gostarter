package mailer

import (
	"github.com/fatkulnurk/gostarter/config"
	"github.com/fatkulnurk/gostarter/pkg/interfaces"
	"github.com/fatkulnurk/gostarter/pkg/logging"
	"github.com/wneessen/go-mail"
)

func NewSmtp(cfg *config.SMTP) (*mail.Client, error) {
	// Deliver the mails via SMTP
	client, err := mail.NewClient(cfg.Host,
		mail.WithSMTPAuth(mail.SMTPAuthType(cfg.AuthType)),
		mail.WithTLSPortPolicy(mail.TLSPolicy(cfg.WithTLSPortPolicy)),
		mail.WithUsername(cfg.Username),
		mail.WithPassword(cfg.Password),
		mail.WithPort(cfg.Port),
	)

	if err != nil {
		logging.Fatalf("Error creating SMTP client: %s", err)
		return nil, err
	}

	return client, nil
}

type SMTPMailer struct {
	client *mail.Client
	from   string
}

func NewSMTPMailer(client *mail.Client, from string) interfaces.IMailer {
	return interfaces.NewSESMailer(client, from)
}
