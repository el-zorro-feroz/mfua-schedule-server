package server

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Printf("Server is listening")
	}()

	grpcServer := grpc.NewServer()

	//TODO: implement proto handlers

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Printf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Printf("ctx.Done: %v", done)
	}

	grpcServer.GracefulStop()
	log.Printf("Server exited")

	return nil
}
