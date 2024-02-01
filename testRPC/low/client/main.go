package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Numbers struct {
	X int
	Y int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	args := Numbers{
		X: 3,
		Y: 5,
	}
	var reply int
	err = client.Call("Numbers.Plus", args, &reply)
	fmt.Println(reply)
}
