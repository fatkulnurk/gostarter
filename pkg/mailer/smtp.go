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
	return &SMTPMailer{
		client: client,
		from:   from,
	}
}

func (s SMTPMailer) SendMail(msg interfaces.MailMessage) error {
	message := mail.NewMsg()
	if err := message.From(s.from); err != nil {
		logging.Fatalf("failed to set FROM address: %s", err)
		return err
	}
	if err := message.To(msg.To); err != nil {
		logging.Fatalf("failed to set TO address: %s", err)
		return err
	}

	message.Subject(msg.Subject)
	if msg.IsHTML {
		message.SetBodyString(mail.TypeTextHTML, msg.Body)
	} else {
		message.SetBodyString(mail.TypeTextPlain, msg.Body)
	}

	if err := s.client.DialAndSend(message); err != nil {
		logging.Fatalf("failed to deliver mail: %s", err)
		return err
	}

	return nil
}
