package main

import "fmt"

func main() {

	ch := make(chan string)
	done := make(chan bool)

outer:
	for {
		go listener(ch)
		select {
		case <-done:
			break outer
		default:
		}
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		ch <- string(i)
	}

	// done <- true

	fmt.Println("main finish")
}

func listener(ch <-chan string) {
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
		}
	}
}
