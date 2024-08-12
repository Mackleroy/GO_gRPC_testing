package main

import (
	"context"
	"second_proto/second_proto"

	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// dial server
	conn, err := grpc.Dial("0.0.0.0:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := second_proto.NewAvgServiceClient(conn)
	stream, err := client.SendNumber(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	log.Println("Client started")

	go func() {
		for i := 1; i <= 30; i++ {
			in := &second_proto.Request{IntValue: int32(i)}
			stream.Send(in)
		}
	}()

	done := make(chan struct{})

	go func() {
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %f", response.GetAvgValue())

		}
	}()

	<-done

}
