package test

import (
	"encoding/json"
	"testing"

	"github.com/streadway/amqp"
	"streamline-post/structs"
)

func TestRabbitMQQueue(t *testing.T) {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		t.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	// Declare a test queue
	q, err := ch.QueueDeclare(
		"test_queue", // name
		false,        // durable
		false,        // autoDelete
		false,        // exclusive
		false,        // noWait
		nil,          // args
	)
	if err != nil {
		t.Fatalf("Failed to declare a queue: %s", err)
	}

	// Publish a test message
	testMessage := structs.PostMessage{
		User: structs.User{
			LinkedInURN: "urn:li:person:123456",
			AccessToken: "test_token",
		},
		Body:       "Test Message",
		Visibility: "PUBLIC",
		Platform:   "linkedin",
	}
	messageBody, err := json.Marshal(testMessage)
	if err != nil {
		t.Fatalf("Failed to marshal test message: %s", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        messageBody,
		})
	if err != nil {
		t.Fatalf("Failed to publish a message: %s", err)
	}
	t.Log("Successfully published a test message to RabbitMQ")
}
