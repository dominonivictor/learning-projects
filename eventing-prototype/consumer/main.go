package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const QUEUE_ENDPOINT = "http://sqs.sa-east-1.localhost.localstack.cloud:4566"
const ACC_ID = "000000000000"
const QUEUE_NAME = "first-queue"

func main() {
	// Connect to LocalStack SQS
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String(QUEUE_ENDPOINT), // LocalStack endpoint
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	svc := sqs.New(sess)
	queueURL := fmt.Sprintf("%s/%s/%s", QUEUE_ENDPOINT, ACC_ID, QUEUE_NAME) // LocalStack uses account ID 000000000000

	// Poll the queue
	for {
		result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: aws.Int64(1),
			WaitTimeSeconds:     aws.Int64(10), // Long polling
		})
		if err != nil {
			log.Printf("Failed to receive message: %v", err)
			continue
		}

		for _, msg := range result.Messages {
			fmt.Printf("Received: %s\n", *msg.Body)
			// Delete the message after processing
			_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				log.Printf("Failed to delete message: %v", err)
			}
		}
	}
}
