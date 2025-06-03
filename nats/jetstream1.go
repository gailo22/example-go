package main

import "github.com/nats-io/nats.go"

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, err := nc.JetStream()

}
