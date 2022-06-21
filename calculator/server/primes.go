package main

import (
	"context"

	pb "github.com/hplxsarthak/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream CalculatorService_PrimesServer) error {

}