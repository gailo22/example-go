package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			fmt.Printf("Goroutine #%d finished.\n", j)
		}(i)
	}
	wg.Wait()
}
