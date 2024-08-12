package second_proto

import (
	"io"
	"log"
)

type avgServiceServer struct {
	UnimplementedAvgServiceServer
}

func NewAvgServiceServer() *avgServiceServer {
	return &avgServiceServer{}
}

func calculate_avg_and_send(list_chan chan []int, srv AvgService_SendNumberServer) {
	for element_list := range list_chan {
		log.Printf("Provided data of values %v", element_list)
		total := 0
		for _, v := range element_list {
			total += v
		}
		avg := float32(total) / float32(len(element_list))
		log.Printf("AVG: %v", avg)
		body := &AvgResponse{AvgValue: avg}
		if err := srv.Send(body); err != nil {
			log.Fatal(err)
		}
	}
}

func (ass *avgServiceServer) SendNumber(srv AvgService_SendNumberServer) error {
	done := make(chan struct{})
	element_list := []int{}
	list_chan := make(chan []int)

	go calculate_avg_and_send(list_chan, srv)

	go func() {
		defer close(list_chan)
		for {
			resp, err := srv.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("error %v", err)
			}
			log.Printf("Received data %v", resp.GetIntValue())
			element_list = append(element_list, int(resp.GetIntValue()))
			if len(element_list) == 10 {
				data := make([]int, 10)
				copy(data, element_list)

				list_chan <- data
				log.Printf("Calculating data of values %v", element_list)

				element_list = []int{}
			}
		}
	}()

	<-done
	return nil
}
