package mailer

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"
)

func joinEmails(emails []string) string {
	var buf bytes.Buffer
	for i, email := range emails {
		if i > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(email)
	}
	return buf.String()
}

func buildRawMessage(ctx context.Context, i InputBuildRawMessage) (*bytes.Buffer, error) {
	// MIME Message
	var rawMessage bytes.Buffer
	if i.Boundary == "" {
		i.Boundary = fmt.Sprintf("MAIN-BOUNDARY-EMAIL-%d", time.Now().Year())
	}

	mainBoundary := i.Boundary
	if i.Destination == nil || i.Destination.ToAddresses == nil {
		return nil, errors.New("destination can't be empty")
	}

	// header from formatnya "Nama <email>"
	fromSender := fmt.Sprintf("%s <%s>", i.Sender.FromName, i.Sender.FromAddress)

	// Header Utama
	rawMessage.WriteString(fmt.Sprintf("From: %s\r\n", fromSender))
	rawMessage.WriteString(fmt.Sprintf("To: %s\r\n", joinEmails(i.Destination.ToAddresses)))
	if i.Destination.CcAddresses != nil {
		rawMessage.WriteString(fmt.Sprintf("Cc: %s\r\n", joinEmails(i.Destination.CcAddresses)))
	}
	rawMessage.WriteString(fmt.Sprintf("Subject: %s\r\n", i.Subject))
	rawMessage.WriteString("MIME-Version: 1.0\r\n")

	// Content-Type berdasarkan jenis konten
	hasBody := i.TextMessage != "" || i.HtmlMessage != ""
	hasAttachments := len(i.Attachments) > 0

	if hasAttachments || (hasBody && (i.TextMessage != "" && i.HtmlMessage != "")) {
		rawMessage.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=\"%s\"\r\n", mainBoundary)) // wajib jika tidak mau dianggap spam
	} else if i.HtmlMessage != "" {
		rawMessage.WriteString("Content-Type: text/html; charset=UTF-8\r\n")
	} else {
		rawMessage.WriteString("Content-Type: text/plain; charset=UTF-8\r\n")
	}
	rawMessage.WriteString("\r\n")

	// Body Content
	if hasBody {
		// Jika ada attachment atau dual content, gunakan boundary
		if hasAttachments || (i.TextMessage != "" && i.HtmlMessage != "") {
			rawMessage.WriteString(fmt.Sprintf("--%s\r\n", mainBoundary))

			// Handle dual content (text + HTML)
			if i.TextMessage != "" && i.HtmlMessage != "" {
				altBoundary := "ALT-BOUNDARY-456" // wajib dikasih boundary
				rawMessage.WriteString(fmt.Sprintf("Content-Type: multipart/alternative; boundary=\"%s\"\r\n\r\n", altBoundary))

				// Text Part
				rawMessage.WriteString(fmt.Sprintf("--%s\r\n", altBoundary))
				rawMessage.WriteString("Content-Type: text/plain; charset=UTF-8\r\n\r\n")
				rawMessage.WriteString(i.TextMessage + "\r\n\r\n")

				// HTML Part
				rawMessage.WriteString(fmt.Sprintf("--%s\r\n", altBoundary))
				rawMessage.WriteString("Content-Type: text/html; charset=UTF-8\r\n\r\n")
				rawMessage.WriteString(i.HtmlMessage + "\r\n\r\n")

				rawMessage.WriteString(fmt.Sprintf("--%s--\r\n\r\n", altBoundary))
			} else {
				// Single content type
				if i.TextMessage != "" {
					rawMessage.WriteString("Content-Type: text/plain; charset=UTF-8\r\n\r\n")
					rawMessage.WriteString(i.TextMessage + "\r\n\r\n")
				} else {
					rawMessage.WriteString("Content-Type: text/html; charset=UTF-8\r\n\r\n")
					rawMessage.WriteString(i.HtmlMessage + "\r\n\r\n")
				}
			}
		} else {
			// Tanpa attachment & single content type
			if i.TextMessage != "" {
				rawMessage.WriteString(i.TextMessage + "\r\n")
			} else {
				rawMessage.WriteString(i.HtmlMessage + "\r\n")
			}
		}
	}

	// Attachments
	for _, attachment := range i.Attachments {
		rawMessage.WriteString(fmt.Sprintf("--%s\r\n", mainBoundary))
		rawMessage.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=\"%s\"\r\n", attachment.Name))
		rawMessage.WriteString(fmt.Sprintf("Content-Type: %s; name=\"%s\"\r\n", attachment.MimeType, attachment.Name))
		rawMessage.WriteString("Content-Transfer-Encoding: base64\r\n\r\n")

		encoder := base64.NewEncoder(base64.StdEncoding, &rawMessage)
		encoder.Write(attachment.Content)
		encoder.Close()
		rawMessage.WriteString("\r\n")
	}

	// Tutup Boundary email
	if hasAttachments || (hasBody && (i.TextMessage != "" && i.HtmlMessage != "")) {
		rawMessage.WriteString(fmt.Sprintf("--%s--\r\n", mainBoundary))
	}

	return &rawMessage, nil
}
