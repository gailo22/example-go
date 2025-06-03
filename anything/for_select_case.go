package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan bool)
	defer close(ch)

	go produce(ch, done)

loop:
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			break loop
		}
	}

	fmt.Println("done")
}

func produce(ch chan int, done chan bool) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	done <- true
}
