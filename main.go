package main

import (
	"log"
	"streamline-post/message_queue"
)

func main() {
	log.Println("Starting Streamline Post Service...")
	message_queue.InitiateQueue()
}