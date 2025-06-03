package main

import "fmt"

func main() {
	world := "world"
	s := []byte(fmt.Sprintf(`{"type": "%v"}`, world))

	fmt.Println(string(s))

	arr := []string{}
	for i := 0; i < 5; i++ {
		arr = append(arr, "xx")
	}

	fmt.Println(arr)
}
