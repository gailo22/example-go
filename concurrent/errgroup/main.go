package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := context.Background()

	wait := errGroup(ctx)

	<-wait

}

func errGroup(ctx context.Context) <-chan struct{} {
	ch := make(chan struct{}, 1)
	var g errgroup.Group

	g.Go(func() error {
		fmt.Println("hello")
		time.Sleep(5 * time.Second)

		for {
			select {
			case <-ctx.Done():
				fmt.Println("context completed")
				return ctx.Err()
			case value, ok := <-ch:
				if !ok {
					return nil
				}
				fmt.Println(value)
			}
		}
	})

	go func() {
		if err := g.Wait(); err != nil {
			fmt.Printf("error: %v", err)
		}
		close(ch)
	}()

	return ch
}
