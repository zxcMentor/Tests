package main

import (
	"log"
	"net"
	"net/rpc"
)

type Summer interface {
	Plus() int
}

type Numbers struct {
	X int
	Y int
}

func (s *Numbers) Plus(args *Numbers, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	sum := new(Numbers)
	rpc.Register(sum)
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server up on port 8080")
	rpc.Accept(l)
}
