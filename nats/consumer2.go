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

	opt, err := nats.NkeyOptionFromSeed("seed.txt")
	if err != nil {
		log.Fatal(err)
	}
	streamName := "BAR"
	subjectName := "BAR.created"

	nc, err := nats.Connect("nats://localhost:4222", opt)
	if err != nil {
		log.Fatal(err)
	}
	js, _ := jetstream.New(nc)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:      streamName,
		Subjects:  []string{fmt.Sprintf("%s.*", streamName)},
		Retention: jetstream.WorkQueuePolicy,
	})

	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		AckPolicy:     jetstream.AckExplicitPolicy,
		FilterSubject: subjectName,
		Durable:       "bar-durable",
	})

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		msgs, _ := consumer.Fetch(2)

		for msg := range msgs.Messages() {
			fmt.Println(string(msg.Data()))
			msg.Ack()
		}

		// msgs, _ := sub.Fetch(10, nats.Context(ctx))
		// for _, msg := range msgs {
		//    msg.Ack()
		//    var order model.Order
		//    err := json.Unmarshal(msg.Data, &order)
		//    if err != nil {
		// 	  log.Fatal(err)
		//    }
		//    log.Println("order-review service")
		//    log.Printf("OrderID:%d, CustomerID: %s, Status:%s\n", order.OrderID, order.CustomerID, order.Status)
		//    reviewOrder(js,order)
		// }
	}

}
