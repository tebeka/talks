package main

import (
	"context"
	"log"
	"time"
)

type Event struct{}

func handleEvent(e Event) {
	time.Sleep(17 * time.Millisecond)
	log.Printf("event")
}

func processEvents(ctx context.Context, ch <-chan Event) {
loop:
	for {
		select {
		case e, ok := <-ch:
			if !ok {
				break loop
			}
			handleEvent(e)
		case <-ctx.Done():
			break loop
		}
	}
}

func main() {
	ch := make(chan Event)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for i := 0; i < 7; i++ {
			ch <- Event{}
		}
		cancel()
	}()

	processEvents(ctx, ch)
}
