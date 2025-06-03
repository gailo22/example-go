package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	consumer(client)

}

func consumer(client pulsar.Client) {
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            "persistent://public/default/my-topic-1",
		SubscriptionName: "my-sub-1",
		Type:             pulsar.Shared,
		Name:             "consumer-1",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer consumer.Close()

	for {
		fmt.Println("receiving message...")
		// may block here
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Received message msgId: %#v -- content: '%s'\n",
			msg.ID().String(), string(msg.Payload()))

		consumer.Ack(msg)
	}

	// if err := consumer.Unsubscribe(); err != nil {
	// 	log.Fatal(err)
	// }
}
