package sms

import "context"

// Request is the struct for sending notice in SMS
type Request struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
	Body string `json:"body,omitempty"`
}

type key int

// SMSRequestKey is the context key for sms.Request content
const SMSRequestKey key = 0

// WithSMSRequest adds the sms request content to the context and return the duplicate
func WithSMSRequest(ctx context.Context, from, to, body string) context.Context {
	return context.WithValue(ctx, SMSRequestKey, &Request{
		From: from,
		Body: body,
		To:   to,
	})
}
