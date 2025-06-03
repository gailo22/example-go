package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("doing work")
		}
	}
}

func main() {
	done := make(chan bool)

	go doWork(done)

	time.Sleep(5 * time.Second)
	close(done)

	// time.Sleep(5 * time.Hour)

}
