package sendgrid

import (
	"context"
	"fmt"
	"log"

	"github.com/azusa0127/notice-hub/email"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendClient is the sendgrid client for sending message out
type SendClient struct {
	c *sendgrid.Client
}

// Send a single message out with sendgrid
func (s *SendClient) Send(ctx context.Context) error {
	r := ctx.Value(email.EmailRequestKey).(*email.Request)
	if r == nil {
		return fmt.Errorf("No EmailRequestKey in the context passed into sendgrid.SendClient.Send")
	}

	message := mail.NewSingleEmail(r.From, r.Subject, r.To, r.PlainText, r.HTML)
	resp, err := s.c.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Body)
	}
	return err
}

// NewSendClient is the alias to sendgrid.NewSendClient
var NewSendClient = sendgrid.NewSendClient
