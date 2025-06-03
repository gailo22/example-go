package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://www.usegolang.com",
		"https://testwithgo.com",
		"https://gophercises.com",
	}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			_, err := http.Get(url)
			if err != nil {
				fmt.Printf("Error pinging: %v\n", url)
				return
			}
			fmt.Printf("Successful ping: %v\n", url)
		}(url)
	}
	wg.Wait()

}
