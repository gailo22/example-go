package main

import (
	"fmt"
	"runtime"
)

func main() {
	emps := 2000
	ch := make(chan string, emps)

	g := runtime.NumCPU()
	sem := make(chan bool, g)

	for e := 0; e < g; e++ {
		go func(emp int) {
			sem <- true
			{
				ch <- "paper"
				fmt.Printf("seceiver: received signal: %d\n", emp)
			}
			<-sem
		}(e)
	}

	for emps > 0 {
		p := <-ch
		emps--
		fmt.Println(p)

	}

	// close(ch)
	fmt.Println("sender: send shutdown signal")

}
