package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "testProj/testGRPC/protos"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	req := &pb.HelloRequest{Name: "Вася", Lastname: "INTRODENSER"}
	clGeo := pb.NewNewGeocodeClient(conn)
	reqGeo := &pb.GeoRequest{
		Lat: "23.32523",
		Lon: "52.13412",
	}
	resGeo, err := clGeo.Geocode(context.Background(), reqGeo)
	if err != nil {
		log.Fatalf("Ошибка при вызове RPC: %v", err)
	}

	log.Printf("Ответ от сервера: %s", resGeo.Message)

	res, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове RPC: %v", err)
	}

	log.Printf("Ответ от сервера: %s", res.Message)
}
