package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	duration := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		ch <- "paper"
	}()

	select {
	case d := <-ch:
		fmt.Println("work completed", d)
	case <-ctx.Done():
		fmt.Println("work cancelled")
	}

}
