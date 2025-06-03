package main

import (
	"encoding/json"
	"log"
	"strconv"

	"model"

	"github.com/nats-io/nats.go"
)

type Order struct {
	OrderID    int
	CustomerID string
	Status     string
}

type Order struct {
	OrderID    int
	CustomerID string
	Status     string
}

const (
	streamName     = "ORDERS"
	streamSubjects = "ORDERS.*"
	subjectName    = "ORDERS.created"
)

func createStream(js nats.JetStreamContext) error {
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Println(err)
	}
	if stream == nil {
		log.Printf("creating stream %q and subjects %q", streamName, streamSubjects)
		js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{streamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func createOrder(js nats.JetStreamContext) error {
	var order Order
	for i := 1; i <= 10; i++ {
		order = model.Order{
			OrderID:    i,
			CustomerID: "Cust-" + strconv.Itoa(i),
			Status:     "created",
		}
		orderJSON, _ := json.Marshal(order)
		_, err := js.Publish(subjectName, orderJSON)
		if err != nil {
			return err
		}
		log.Printf("Order with OrderID:%d has been published\n", i)
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Connect to NATS
	nc, _ := nats.Connect(nats.DefaultURL)
	// Creates JetStreamContext
	js, err := nc.JetStream()
	checkErr(err)
	// Creates stream
	err = createStream(js)
	checkErr(err)
	// Create orders by publishing messages
	err = createOrder(js)
	checkErr(err)
}
