package mailer

import "context"

type IMailer interface {
	SendMail(ctx context.Context, msg InputSendMail) (*OutputSendMail, error)
}

type Sender struct {
	// Override Sender's email address, which may be optional
	FromAddress string
	// Override Sender's name, which may be optional
	FromName string
}

type Destination struct {
	// An array that contains the email addresses of the "To" recipients for the email.
	ToAddresses []string

	// An array that contains the email addresses of the "BCC" (blind carbon copy)
	// recipients for the email.
	BccAddresses []string

	// An array that contains the email addresses of the "CC" (carbon copy) recipients
	// for the email.
	CcAddresses []string
}

type Attachment struct {
	// The content of the attachment in []byte.
	Content []byte

	// The name of the attachment. eg: example.pdf
	Name string

	// The content type of the attachment.
	// For example, "application/pdf" or "application/xml".
	MimeType string
}

type InputSendMail struct {
	Subject     string
	TextMessage string
	HtmlMessage string
	Destination Destination
	Attachments []Attachment
	// Boundary for MIME parts (optional) , if set "", system will generate boundary
	Boundary string
	// Override Sender's email address, which may be optional
	Sender *Sender
}

type OutputSendMail struct {
	MessageID *string
}

type InputBuildRawMessage struct {
	// Email subject
	Subject string

	// Email's plain text content, with header text/plain; charset="UTF-8"
	TextMessage string

	// Email's HTML content, with header text/html; charset="UTF-8"
	HtmlMessage string

	// Sender using for sent email, required
	Sender Sender

	// Email's recipient details
	Destination *Destination

	// Attachments to be added to the email
	Attachments []Attachment

	// Boundary for MIME parts
	Boundary string
}
