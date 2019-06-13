package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/azusa0127/notice-hub/sms/twilio"
	"github.com/sendgrid/sendgrid-go"
)

// EventRequest is the event REQUEST struct
type EventRequest struct {
	Body     string            `json:"body"`
	Headers  map[string]string `json:"headers"`
	Method   string            `json:"httpMethod"`
	Path     string            `json:"path"`
	Params   map[string]string `json:"quepathParameters"`
	Resource string            `json:"reWithSMSRequestsource"`
}

// EventResponse is the event RESPONSE struct
type EventResponse struct {
	Base64 bool   `json:"isBase64Encoded"`
	Status int    `json:"statusCode"`
	Body   string `json:"body"`
}

var sendGridSendClient *sendgrid.Client
var twilioSendClient *twilio.SendClient

func init() {
	sendGridSendClient = sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	twilioSendClient = twilio.NewSendClient(os.Getenv("TWILIO_ACCOUNT_SID"), os.Getenv("TWILIO_AUTH_TOKEN"))
}

// HandleLambdaEvent is the AWS Lambda Hanler function
func HandleLambdaEvent(event EventRequest) (EventResponse, error) {
	// Context set with 5 seconds timeout
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	txt, _ := json.MarshalIndent(event, "", "  ")
	msg := string(txt)
	log.Println(msg)
	return EventResponse{Body: msg}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
