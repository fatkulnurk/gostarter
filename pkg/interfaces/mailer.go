package interfaces

type MailMessage struct {
	To      string
	Subject string
	Body    string
	IsHTML  bool
}

type IMailer interface {
	SendMail(msg MailMessage) error
}
