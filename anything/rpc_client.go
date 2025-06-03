package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	serverAddress := "127.0.0.1"

	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}

	args := Args{3, 4}
	var reply int

	client.Call("MultiplyService.Do", args, &reply)
	log.Printf(" %d*%d=%d", args.A, args.B, reply)

}
