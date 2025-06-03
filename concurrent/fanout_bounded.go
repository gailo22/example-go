package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	work := []string{"paper", "paper", 2000: "paper"}

	g := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(g)

	ch := make(chan string, g)

	for e := 0; e < g; e++ {
		go func(emp int) {
			defer wg.Done()
			for p := range ch {
				fmt.Printf("seceiver: received signal: %d, %s\n", emp, p)
			}
			fmt.Println("seceiver: received shutdown signal")
		}(e)
	}

	for _, wrk := range work {
		ch <- wrk
	}

	close(ch)
	wg.Wait()

}
