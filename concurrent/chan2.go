package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Read from channel
		fmt.Println("value:", <-ch)
	}()

	// Send to channel
	// ch <- 42

	wg.Wait()
	fmt.Println("Program completed")
}
