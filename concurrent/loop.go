package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func mainloop() {
	exitSignal := make(chan os.Signal, 1)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM)
	<-exitSignal

	systemTeardown()
}

func main() {
	systemStart()
	listener1()
	mainloop()
}

func systemStart() {
	fmt.Println("system start...")
}

func listener1() {
	ch := make(chan int)

	go func() {
		for {
			select {
			case value := <-ch:
				fmt.Println("value:", value)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second)
			ch <- 42
		}
	}()
}

func systemTeardown() {
	fmt.Println("system tear down!")
}
