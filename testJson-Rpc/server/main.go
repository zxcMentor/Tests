package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type MathService struct{}

type Args struct {
	Query string `json:"query"`
}

type Address struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	City     string `json:"city"`
}

func (m *MathService) Plus(args *Args, reply *[]*Address) error {
	*reply = []*Address{{
		Name:     args.Query,
		LastName: args.Query,
		City:     args.Query,
	}}
	return nil
}

func main() {
	mathService := new(MathService)
	rpc.Register(mathService)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	defer l.Close()

	log.Println("Serving RPC server on port 1234")
	http.Serve(l, nil)
}
