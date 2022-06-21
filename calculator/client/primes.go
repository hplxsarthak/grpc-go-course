package main

import (
	"context"
	"io"
	"log"

	pb "github.com/hplxsarthak/grpc-go-course/calculator/proto"
)

func doPrimes (c pb.CalculatorServiceClient) {
	log.Printf("doPrimes was invoked")

	req := &pb.PrimeRequest{
		Number: 15,
	}

	stream,err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while callling Primes: %v\n", err)
	}

	for {
		res,err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Fatalf("Primes: %d\n", res.Result)
	}
	
}