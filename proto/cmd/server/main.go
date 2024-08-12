package main

import (
	"first_proto/first_proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		log.Fatalf("Failed to listen grpc %v", err)
	}

	s := grpc.NewServer()
	first_proto.RegisterTimeServiceServer(s, first_proto.NewTimeServiceServerCustom())

	log.Println("Server starts running")
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stoped running")
}
