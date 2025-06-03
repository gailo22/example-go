package main

import "fmt"

func main() {
	const cap = 100
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Println("received:", p)
		}
	}()

	const work = 2000
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("sending: sent", w)
		default:
			fmt.Println("sending: drop data", w)
		}
	}

	close(ch)
}
