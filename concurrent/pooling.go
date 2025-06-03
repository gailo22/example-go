package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan string)

	g := runtime.NumCPU()
	for e := 0; e < g; e++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("seceiver: received signal: %d, %s\n", emp, p)
			}
			fmt.Println("seceiver: received shutdown signal")
		}(e)
	}

	const work = 100
	for w := 0; w < work; w++ {
		ch <- "work"
		fmt.Println("sender: send signal:", w)
	}

	close(ch)
	fmt.Println("sender: send shutdown signal")

}
