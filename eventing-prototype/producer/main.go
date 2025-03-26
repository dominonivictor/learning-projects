package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type PriceUpdate struct {
	Asset string `json:"asset"`

	PriceInMiliDollars int64 `json:"priceInMiliDollars"` // 0.001 = 1 Mili Dollar
}

const QUEUE_ENDPOINT = "http://sqs.sa-east-1.localhost.localstack.cloud:4566"
const ACC_ID = "000000000000"
const QUEUE_NAME = "first-queue"

func main() {
	// Connect to LocalStack SQS (running locally)
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String(QUEUE_ENDPOINT), // LocalStack endpoint
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	svc := sqs.New(sess)

	queueURL := fmt.Sprintf("%s/%s/%s", QUEUE_ENDPOINT, ACC_ID, QUEUE_NAME) // LocalStack uses account ID 000000000000

	// Simulate sending price updates
	for {
		plus := int64(rand.Intn(2000))
		minus := int64(rand.Intn(2000))
		update := PriceUpdate{Asset: "BTC", PriceInMiliDollars: 30000 + plus - minus}
		body, _ := json.Marshal(update)

		_, err := svc.SendMessage(&sqs.SendMessageInput{
			QueueUrl:    aws.String(queueURL),
			MessageBody: aws.String(string(body)),
		})
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		} else {
			fmt.Printf("Sent: %s\n", body)
		}
		time.Sleep(100 * time.Millisecond)
	}
}
