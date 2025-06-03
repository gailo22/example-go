package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
)

func main() {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:               "pulsar://localhost:6650",
		OperationTimeout:  30 * time.Second,
		ConnectionTimeout: 30 * time.Second,
	})
	if err != nil {
		log.Fatalf("Could not instantiate Pulsar client: %v", err)
	}

	defer client.Close()

	producer(client)

}

func producer(client pulsar.Client) {

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic: "persistent://public/default/my-topic-1",
	})

	if err != nil {
		log.Fatal(err)
	}

	for {

		_, err = producer.Send(context.Background(), &pulsar.ProducerMessage{
			Payload: []byte(`{"deviceId":"TDAD9JOPIJGJLOZ7","aliasName":"TDAD9JOPIJGJLOZ7","title":"Motion alert","body":"Person may be detected from TDAD9JOPIJGJLOZ7","type":5010,"cognitives":"humanoid","imageUrl":"https://td-alarm-apse2.obs.ap-southeast-2.myhuaweicloud.com/20240830/cdf4d8558b6fca44c214827c33894537/20240821.i.1724242305.94b49d4e-2640-4578-521b-2111ec69fb4e.jpg?expires=1724847105&stor=obs","ts":1724242305000}`),
		})

		time.Sleep(1 * time.Second)
	}

	defer producer.Close()

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")

}
