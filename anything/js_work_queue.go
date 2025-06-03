package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	nc, _ := nats.Connect("localhost:4222")
	defer nc.Drain()

	js, _ := jetstream.New(nc)

	cfg := jetstream.StreamConfig{
		Name:      "EVENTS",
		Retention: jetstream.WorkQueuePolicy,
		Subjects:  []string{"EVENTS.*"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := js.CreateStream(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("created the stream")

	// js.Publish(ctx, "EVENTS.created", []byte("msg 1"), nil)
	// js.Publish(ctx, "EVENTS.created", []byte("msg 2"), nil)
	// js.Publish(ctx, "EVENTS.created", []byte("msg 3"), nil)
	// fmt.Println("published 3 messages")

	for i := 0; i < 10; i++ {
		_, err = js.PublishMsg(ctx, &nats.Msg{
			Data:    []byte(fmt.Sprintf("hello %d", i)),
			Subject: "EVENTS.created",
		})
	}

	fmt.Println("# Stream info without any consumers")
	printStreamState(ctx, stream)

	cons1, _ := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		AckPolicy:     jetstream.AckExplicitPolicy,
		Name:          "processor-1",
		FilterSubject: "EVENTS.created",
	})

	iter, _ := cons1.Messages(jetstream.PullMaxMessages(1))
	numWorkers := 5
	sem := make(chan struct{}, numWorkers)
	for {
		sem <- struct{}{}
		go func() {
			defer func() {
				<-sem
			}()
			msg, err := iter.Next()
			if err != nil {
				// handle err
			}
			fmt.Printf("Processing msg: %s\n", string(msg.Data()))
			doWork(msg)
			msg.Ack()
		}()
	}

	// msgs, _ := cons1.Fetch(3)
	// for msg := range msgs.Messages() {
	// 	fmt.Printf("cons2 sub got: %s\n", msg.Subject())
	// 	msg.DoubleAck(ctx)
	// }

	// fmt.Println("\n# Stream info with one consumer")
	// printStreamState(ctx, stream)

	// cons2, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
	// 	Name:          "processor-2",
	// 	FilterSubject: "events.created",
	// })
	// log.Fatal(err)

	// msgs, _ := cons1.Fetch(2)
	// for msg := range msgs.Messages() {
	// 	fmt.Printf("cons1 sub got: %s\n", msg.Subject())
	// 	msg.Ack()
	// }

	// msgs, _ = cons2.Fetch(2)
	// for msg := range msgs.Messages() {
	// 	fmt.Printf("cons2 sub got: %s\n", msg.Subject())
	// 	msg.Ack()
	// }

}

func doWork(msg jetstream.Msg) {
	fmt.Println(string(msg.Data()))
}

func printStreamState(ctx context.Context, stream jetstream.Stream) {
	info, _ := stream.Info(ctx)
	b, _ := json.MarshalIndent(info.State, "", " ")
	fmt.Println(string(b))
}
