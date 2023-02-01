package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	log.Println("application started")

	<-ctx.Done()

	log.Println("application stopped")
}
