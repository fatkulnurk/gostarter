package mailer

import (
	"github.com/fatkulnurk/gostarter/config"
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
