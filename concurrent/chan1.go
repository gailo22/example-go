package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	exit := make(chan bool, 1)

	go func() {
		for i := 0; i < 2; i++ {
			ch <- i
		}
		close(exit)
	}()

	fmt.Println("value:", <-ch)
	fmt.Println("value:", <-ch)
	// fmt.Println("value:", <-ch)
	<-exit
}
