package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	const maxConcurrency = 3
	sem := make(chan struct{}, maxConcurrency) // Semaphore with 3 slots

	var wg sync.WaitGroup
	msgChan := getMessages1()

	for msg := range msgChan {
		sem <- struct{}{} // Acquire semaphore
		wg.Add(1)

		go func(m string) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore

			processMessage1(m)
		}(msg)
	}

	wg.Wait()
}

func getMessages1() <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- fmt.Sprintf("Message %d", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	return ch
}

func processMessage1(msg string) {
	delay := time.Duration(500+rand.Intn(1000)) * time.Millisecond
	fmt.Printf("Processing: %s (will take %v)\n", msg, delay)
	time.Sleep(delay)
	fmt.Printf("Finished: %s\n", msg)
}
