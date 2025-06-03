package main

import (
	"fmt"
	"slices"
)

func main() {

	things := []string{"foo", "bar", "baz"}
	fmt.Println(slices.Contains(things, "foo"))
	fmt.Println(slices.Contains(things, "has"))

	key1 := `"abcdep"`
	key2 := key1[1 : len(key1)-1]
	fmt.Println("key1:", key1)
	fmt.Println("key2:", key2)
}
