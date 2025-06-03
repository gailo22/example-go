package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numMsgs := 10
	numWorkers := 2
	var wg sync.WaitGroup

	workQueue := make(chan int)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()
			for msgId := range workQueue { // block
				fmt.Println("receiving:", msgId)
				slowRPC(workerId, msgId)
			}
		}(i)
	}

	for msgId := 0; msgId < numMsgs; msgId++ {
		fmt.Println("sending:", msgId)
		workQueue <- msgId //block
	}

	close(workQueue)

	wg.Wait()
}

func slowRPC(workerId, msgId int) {
	time.Sleep(1 * time.Second)
	fmt.Printf("slow rpc workerId: %d, msgId: %d\n", workerId, msgId)
}
