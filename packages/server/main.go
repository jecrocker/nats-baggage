package main

import (
	"github.com/jecrocker/nats-baggage/packages/server/logging"
	"github.com/jecrocker/nats-baggage/packages/server/storage"
)

func main() {
	logger := logging.New("main")

	baggage := storage.New()

	logger.Infow("Information")

	baggage.Store("hello", "world")
	baggage.Get("hello", "2021")
}
