package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	maxConcurrency = 3
	maxRetries     = 3
	timeoutPerTask = 1 * time.Second
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	sem := make(chan struct{}, maxConcurrency)
	var wg sync.WaitGroup
	msgChan := getMessages2()

	for msg := range msgChan {
		sem <- struct{}{}
		wg.Add(1)

		go func(m string) {
			defer wg.Done()
			defer func() { <-sem }()

			success := false
			for attempt := 1; attempt <= maxRetries; attempt++ {
				ctx, cancel := context.WithTimeout(context.Background(), timeoutPerTask)
				err := processMessageWithContext(ctx, m, attempt)
				cancel()

				if err == nil {
					success = true
					break
				}

				fmt.Printf("Retry %d for %s due to error: %v\n", attempt, m, err)
				time.Sleep(200 * time.Millisecond) // backoff between retries
			}

			if !success {
				fmt.Printf("Failed to process %s after %d attempts\n", m, maxRetries)
			}
		}(msg)
	}

	wg.Wait()
}

func getMessages2() <-chan string {
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

func processMessageWithContext(ctx context.Context, msg string, attempt int) error {
	delay := time.Duration(500+rand.Intn(1000)) * time.Millisecond // 500ms–1500ms
	fmt.Printf("[Attempt %d] Processing: %s (will take %v)\n", attempt, msg, delay)

	select {
	case <-time.After(delay):
		fmt.Printf("[Attempt %d] Finished: %s\n", attempt, msg)
		return nil
	case <-ctx.Done():
		return fmt.Errorf("timeout")
	}
}
