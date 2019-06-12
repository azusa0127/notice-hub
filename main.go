package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// MyEvent is the event REQUEST struct
type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

// MyResponse is the event RESPONSE struct
type MyResponse struct {
	Message string `json:"Answer"`
	Count   int    `json:"Invocation count"`
}

var invokeCount = 0
var initialized = false

func init() {
	initialized = true
}

// HandleLambdaEvent is the AWS Lambda Hanler function
func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	if initialized {
		log.Println("Hanlder is initialized")
	}
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age), Count: invokeCount}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
