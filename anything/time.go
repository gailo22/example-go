package main

import (
	"fmt"
	"time"
)

func main() {
	currentDt := time.Now().UTC()

	time.Sleep(1 * time.Second)

	processTimes := time.Since(currentDt).Milliseconds()
	fmt.Println(processTimes)
}
