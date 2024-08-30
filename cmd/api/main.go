package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hhertout/grpc_boilerplate/internal/interceptor"
	"github.com/hhertout/grpc_boilerplate/internal/server"
	"github.com/hhertout/grpc_boilerplate/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Failed to create listener")
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(interceptor.GetInterceptor()...),
	)
	reflection.Register(s)
	defer s.GracefulStop()

	pb.RegisterCalculatorServer(s, &server.Server{})

	fmt.Println("Lauching server")
	if err := s.Serve(listener); err != nil {
		s.GracefulStop()
		log.Fatalf("Fail to serve : %v", err)
	}
}
