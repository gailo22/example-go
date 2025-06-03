package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	opt, err := nats.NkeyOptionFromSeed("../seed.txt")
	if err != nil {
		log.Fatal(err)
	}
	nc, err := nats.Connect("nats://localhost:4222", opt)
	if err != nil {
		log.Fatal(err)
	}

	// js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// js.AddStream(&nats.StreamConfig{
	// 	Name:     "ORDERS",
	// 	Subjects: []string{"ORDERS.processed"},
	// })

	// sub, _ := js.PullSubscribe("ORDERS.processed", "xx", nats.PullMaxWaiting(128), nats.MaxDeliver(2))

	js, _ := jetstream.New(nc)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	stream, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.processed"},
	})
	sub, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: "ORDERS.*",
		Durable:       "xxx",
		MaxDeliver:    3,
	})

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		msgs, _ := sub.Fetch(10)
		for msg := range msgs.Messages() {
			fmt.Println(string(msg.Data()))

			msg.NakWithDelay(time.Second * 5)
			tryAgain, apiRespose := callAPIWithDelay()
			fmt.Println("apiResponse:", apiRespose)
			if tryAgain {
				msg.NakWithDelay(time.Second * 5)
				continue
			}
			msg.Ack()
		}

		fmt.Println("...")
	}
}

func callAPIWithDelay() (bool, string) {
	// rand.Seed(time.Now().UnixNano())
	// n := rand.Intn(10) // n will be between 0 and 10
	// if n < 5 {
	// 	return true, ""
	// }
	// time.Sleep(time.Duration(n) * time.Second)
	// return false, fmt.Sprintf("done with %v seconds delay", n)
	return true, "xxx"
}
