package controller

import (
	"log"
	"net/http"
	"net/rpc"
	"strconv"
)

type MathHandler interface {
	PlusHandler(w http.ResponseWriter, r *http.Request)
}

type MathHandle struct {
	ClientRpc *rpc.Client
}

func NewMathHand() *MathHandle {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Println("dialing:", err)
	}
	return &MathHandle{client}
}

func (m *MathHandle) PlusHandler(w http.ResponseWriter, r *http.Request) {

	args := []int{2, 5}
	var reply int
	err := m.ClientRpc.Call("MathService.Plus", args, &reply)
	if err != nil {
		log.Fatal("err:", err)
	}
	res := strconv.Itoa(reply)
	w.Write([]byte(res))
	log.Println("handle ok")
}
