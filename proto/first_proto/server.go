package first_proto

import (
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

type MyError struct{}

func (m *MyError) Error() string {
	return "Duration must be provided, it has to be positive value"
}

type timeServiceServerCustom struct {
	UnimplementedTimeServiceServer
}

func NewTimeServiceServerCustom() *timeServiceServerCustom {
	return &timeServiceServerCustom{}
}

func (tsc *timeServiceServerCustom) StreamTime(in *Request, srv TimeService_StreamTimeServer) error {
	duration := in.GetDurationSecs()
	if duration == 0 {
		return &MyError{}
	}

	log.Println("Duration: ", duration)

	go func(duration uint32) {
		ticker := time.Tick(time.Duration(duration) * time.Second)
		for range ticker {
			response := TimeResponse{CurrentTime: timestamppb.Now()}

			if err := srv.Send(&response); err != nil {
				log.Printf("send error %v", err)
			}
		}

	}(duration)

	var stopper string
	fmt.Scan(&stopper)
	return nil
}
