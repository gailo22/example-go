package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

type Message struct {
	ID        int
	Type      string
	Body      UserAction
	Timestamp time.Time
}

type UserAction struct {
	UserID    int
	Action    string
	Succeeded bool
	Error     string
	Timestamp time.Time
}

func main() {

	fmt.Println("Connection to NATS...")
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)

	nc.Subscribe("hello.world", handleMessage)
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	go http.ListenAndServe(":8080", nil)

	runtime.Goexit()

}

func handleMessage(m *nats.Msg) {
	fmt.Println(m)
}
