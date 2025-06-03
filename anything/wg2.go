package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	// Let Go use all available CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano()) // Seed the random generator

	var wg sync.WaitGroup
	msgChan := getMessages()

	for msg := range msgChan {
		wg.Add(1)
		go func(m string) {
			defer wg.Done()
			processMessage(m)
		}(msg)
	}

	wg.Wait()
}

func getMessages() <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- fmt.Sprintf("Message %d", i)
			// Simulate fast incoming messages
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return ch
}

func processMessage(msg string) {
	delay := time.Duration(500+rand.Intn(1000)) * time.Millisecond // 500ms to 1500ms
	fmt.Printf("Processing: %s (will take %v)\n", msg, delay)
	time.Sleep(delay)
	fmt.Printf("Finished: %s\n", msg)
}
