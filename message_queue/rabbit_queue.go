package message_queue

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"streamline-post/structs"
	"streamline-post/post"
)

func InitiateQueue() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err) 
		return
	}
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err) 
		return
	}
	defer func() {
		if ch != nil {
			ch.Close()
		}
	}()

	q, err := ch.QueueDeclare(
		"rabbit_queue", // name
		false,          // durable
		false,          // autoDelete
		false,          // exclusive
		false,          // noWait
		nil,            // args
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err) 
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // autoAck
		false,  // exclusive
		false,  // noLocal
		false,  // noWait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err) 
		return
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			handlePost(d.Body)
		}
	}()

	<-forever
}

func handlePost(message []byte) {
	var post_struct structs.PostMessage
	err := json.Unmarshal(message, &post_struct)
	if err != nil {
		log.Printf("Error parsing message: %s", err)
		return
	}

	switch post_struct.Platform {
	case "linkedin":
		err = post.LinkedInPost(post_struct)
	case "x":
		err = post.XPost(post_struct)
	default:
		log.Println("Received unknown platform")
	}
}