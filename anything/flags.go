package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "John", "This is the name")

	flag.Parse()

	fmt.Println("name: ", *name)
}
