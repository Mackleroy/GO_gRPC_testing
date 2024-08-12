package main

import (
	"context"
	"first_proto/first_proto"
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
	client := first_proto.NewTimeServiceClient(conn)
	in := &first_proto.Request{DurationSecs: 2}
	stream, err := client.StreamTime(context.Background(), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	log.Println("Client started")

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
			log.Printf("Resp received: %s", response.GetCurrentTime())

		}
	}()

	<-done

}
