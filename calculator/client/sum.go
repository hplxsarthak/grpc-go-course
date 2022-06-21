package main

import (
	"context"
	"log"

	pb "github.com/hplxsarthak/grpc-go-course/calculator/proto"
)

func doSum( c pb.CalculatorServiceClient){
	log.Println("doCalculator was invoked")
	res,err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber: 2,
		SecondNumber: 3,
	})

	if err != nil {
		log.Fatalf("Could not sum %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}