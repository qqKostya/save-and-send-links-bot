package event_consumer

import (
	"log"
	"time"

	"example.com/m/events"
)

type Consumer struct {
	fetcher    events.Fetcher
	processor  events.Processor
	batchStize int
}

func New(fetcher events.Fetcher, processor events.Processor, batchStize int) Consumer {
	return Consumer{
		fetcher:    fetcher,
		processor:  processor,
		batchStize: batchStize,
	}
}

func (c Consumer) Start() error {
	for {
		gotEvent, err := c.fetcher.Fetch(c.batchStize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvent) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(gotEvent); err != nil {
			log.Print(err)

			continue
		}
	}
}

func (c Consumer) handleEvents(events []events.Event) error {
	for _, event := range events {
		log.Printf("got new event", event.Text)

		if err := c.processor.Process(event); err != nil {
			log.Printf("can't handle event: %s", err.Error())

			continue
		}
	}

	return nil
}
