package main

import (
	"flag"
	"log"

	tgClient "example.com/m/clients/telegram"
	event_consumer "example.com/m/consumer/event-consumer"
	"example.com/m/events/telegram"
	"example.com/m/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	tgClient := tgClient.New(tgBotHost, mustToken())

	eventProcessor := telegram.New(&tgClient, files.New(storagePath))

	log.Print("service started")

	consumer := event_consumer.New(eventProcessor, eventProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
