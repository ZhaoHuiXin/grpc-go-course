package main

import (
	"google.golang.org/grpc"
	"net"
	"github.com/prometheus/common/log"
	"grpc-go-course/calculator/calculatorpb"
	"context"
	"fmt"
	"time"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error){
	fmt.Printf("Received Sum RPC: %v\n", req)
	n1 := req.GetFirstNumber()
	n2 := req.GetSecondNumber()
	result := n1 + n2
	res := &calculatorpb.SumResponse{
		SumResult: result,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest,
	stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error{
		// or directly call req.PrimeNumber
	primeNumber := req.GetPrimeNumber()
	var k int32 = 2
	for primeNumber > 1{
		if primeNumber % k == 0{
			result := &calculatorpb.PrimeNumberDecompositionResponse{
				Result: k,
			}
			fmt.Println(k)
			err := stream.Send(result)
			time.Sleep(3 * time.Second)
			if err != nil{
				log.Fatalf("error while send result %v", err)
			}
			primeNumber = primeNumber / k
		}else{
			k += 1
		}
	}
	return nil
}

func main(){
	fmt.Println("begin calculator service.")
	listen, err := net.Listen("tcp", "localhost:50052")
	if err != nil{
		log.Fatalf("error: %v", err)
	}
	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	s.Serve(listen)
}
