package telegram

import "example.com/m/clients/telegram"

type Processor struct {
	tg     *telegram.Client
	offset int
	// storage
}

func New(client *telegram.Client) {}
