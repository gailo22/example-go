package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type MultiplyService struct{}

func (t *MultiplyService) Do(args *Args, reply *int) error {
	log.Println("inside MuliplyService")
	*reply = args.A * args.B
	return nil
}

func main() {
	service := new(MultiplyService)
	rpc.Register(service)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	http.Serve(l, nil)

}
