package main

import (
	"log"
	"net"
	"second_proto/second_proto"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatalf("Failed to listen grpc %v", err)
	}

	s := grpc.NewServer()
	second_proto.RegisterAvgServiceServer(s, second_proto.NewAvgServiceServer())

	log.Println("Server starts running")
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped running")
}
