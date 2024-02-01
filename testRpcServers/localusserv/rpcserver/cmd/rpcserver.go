package main

import (
	"log"
	"net"
	"net/rpc"
)

type MathServicer interface {
	Plus(args []int, reply *int) error
}

type MathService struct {
}

func (m *MathService) Plus(args []int, reply *int) error {
	*reply = args[0] + args[1]
	return nil
}

func main() {
	math := new(MathService)
	rpc.Register(math)

	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Ошибка при запуске сервера:", err)
	}

	log.Println("Сервер запущен на порту 1234")
	rpc.Accept(l)
}
