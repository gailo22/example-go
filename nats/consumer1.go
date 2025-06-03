package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	opt, err := nats.NkeyOptionFromSeed("seed.txt")
	if err != nil {
		log.Fatal(err)
	}
	nc, err := nats.Connect("nats://localhost:4222", opt)
	if err != nil {
		log.Fatal(err)
	}

	// Use the JetStream context to produce and consumer messages
	// that have been persisted.
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	js.AddStream(&nats.StreamConfig{
		Name:     "FOO",
		Subjects: []string{"FOO.created"},
	})

	js.Publish("foo", []byte("Hello JS!"))

	// Publish messages asynchronously.
	// for i := 0; i < 500; i++ {
	// 	js.PublishAsync("foo", []byte("Hello JS Async!"))
	// }
	// select {
	// case <-js.PublishAsyncComplete():
	// case <-time.After(5 * time.Second):
	// 	fmt.Println("Did not resolve in time")
	// }

	// Create Pull based consumer with maximum 128 inflight.
	sub, _ := js.PullSubscribe("FOO.created", "wq", nats.PullMaxWaiting(128))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		// Fetch will return as soon as any message is available rather than wait until the full batch size is available, using a batch size of more than 1 allows for higher throughput when needed.
		msgs, _ := sub.Fetch(10, nats.Context(ctx))
		for _, msg := range msgs {
			fmt.Println(string(msg.Data))
			msg.Ack()
		}
	}
}
