package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	Query string `json:"query"`
}

type Address struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	City     string `json:"city"`
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &Args{
		Query: "your search query",
	}
	var reply []*Address
	err = client.Call("MathService.Plus", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}

	log.Printf("Reply: %+v\n", reply)
}
