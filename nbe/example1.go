package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

type payload struct {
	Foo string `json:"foo"`
	Bar int    `json:"bar"`
}

func main() {
	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)
	defer nc.Drain()

	sub, _ := nc.SubscribeSync("foo")
	sub.AutoUnsubscribe(2)

	p := &payload{
		Foo: "bar",
		Bar: 27,
	}

	p_json, _ := json.Marshal(p)

	nc.Publish("foo", p_json)
	nc.Publish("foo", []byte("not json"))

	var dat payload
	for {
		msg, err := sub.NextMsg(time.Second)
		if err != nil {
			break
		}

		err = json.Unmarshal(msg.Data, &dat)
		if err != nil {
			fmt.Printf("received invalid JSON payload: %s\n", msg.Data)
		} else {
			fmt.Printf("received valid JSON payload: %+v\n", dat)
		}

	}

}
