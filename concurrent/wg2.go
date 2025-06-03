package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // 1

func routine() {
	defer wg.Done() // 3
	time.Sleep(3 * time.Second)
	fmt.Println("routine finished")
}

func main() {
	wg.Add(1)    // 2
	go routine() // *
	wg.Wait()    // 4
	fmt.Println("main finished")
}
