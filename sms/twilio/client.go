package twilio

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/azusa0127/notice-hub/sms"
)

type transportDecorated struct {
	rt         http.RoundTripper
	accountSid string
	authToken  string
}

func (t *transportDecorated) RoundTrip(r *http.Request) (*http.Response, error) {
	r.URL.User = url.UserPassword(t.accountSid, t.authToken)
	return t.rt.RoundTrip(r)
}

// SendClient is the twilio client for sending message out
type SendClient struct {
	c          *http.Client
	accountSid string
}

// NewSendClient create a SendClient corresponding to specific authentication
func NewSendClient(accountSid string, authToken string) *SendClient {
	return &SendClient{
		c: &http.Client{
			Transport: &transportDecorated{
				rt:         http.DefaultTransport,
				accountSid: accountSid,
				authToken:  authToken,
			},
		},
	}
}

// Send a single message out with twilio
func (s *SendClient) Send(ctx context.Context) error {
	r := ctx.Value(sms.SMSRequestKey).(*sms.Request)
	if r == nil {
		return fmt.Errorf("No SMSRequestKey in the context passed into twilio.SendClient.Send")
	}
	data := url.Values{}
	data.Set("From", r.From)
	data.Set("Body", r.Body)
	data.Set("To", r.To)
	resp, err := s.c.Post("https://api.twilio.com/2010-04-01/Accounts/"+s.accountSid+"/Messages.json", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Body)
	}
	return err
}
