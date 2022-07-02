package main

import (
	"context"
	"log"
	"time"

	pb "github.com/hplxsarthak/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadlines(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", in)

	for i:=0; i<3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client cancelled the request!")
			return nil, status.Error(codes.Canceled,"The client cancelled the request!")
		}

		time.Sleep(1 * time.Second)
	}

	// if deadline is not exceeded we reach here
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	},nil
}