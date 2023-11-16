package main

import (
	"log"
	"schedule/internal/server"
)

func main() {
	server := server.NewServer()
	if err := server.Run(); err != nil {
		log.Printf("Failed to start server. Reason: %v", err)
	}
}
