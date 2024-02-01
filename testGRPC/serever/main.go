package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "testProj/testGRPC/protos"
)

type GreeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := "ZDAROVA, " + req.Name + " " + req.Lastname + "!"
	return &pb.HelloResponse{Message: message}, nil
}

type GeoServer struct {
	pb.UnimplementedNewGeocodeServer
}

func (g *GeoServer) Geocode(ctx context.Context, req *pb.GeoRequest) (*pb.GeoResponse, error) {
	address := gofakeit.Address()

	message := fmt.Sprintf("%f and %f this is: city: %s, address: %s", address.Latitude, address.Longitude, address.City, address.Address)
	return &pb.GeoResponse{Message: message}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при прослушивании порта: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	pb.RegisterNewGeocodeServer(server, &GeoServer{})
	log.Println("Запуск gRPC сервера...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
