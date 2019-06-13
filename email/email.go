package email

import (
	"context"

	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Request is the struct for sending notice in email
type Request struct {
	From      *mail.Email `json:"from,omitempty"`
	To        *mail.Email `json:"to,omitempty"`
	Subject   string      `json:"subject,omitempty"`
	PlainText string      `json:"plainTextContent,omitempty"`
	HTML      string      `json:"htmlContent,omitempty"`
}

type key int

// EmailRequestKey is the context key for sms.Request content
const EmailRequestKey key = 0

// WithSMSRequest adds the sms request content to the context and return the duplicate
func WithSMSRequest(ctx context.Context, from, to *mail.Email, subject, plainText, html string) context.Context {
	return context.WithValue(ctx, EmailRequestKey, &Request{
		From:      from,
		To:        to,
		Subject:   subject,
		PlainText: plainText,
		HTML:      html,
	})
}
