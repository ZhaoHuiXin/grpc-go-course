package main

import (
	"google.golang.org/grpc"
	"net"
	"github.com/prometheus/common/log"
	"grpc-go-course/calculator/calculatorpb"
	"context"
	"fmt"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error){
	n1 := req.GetFirstNumber()
	n2 := req.GetSecondNumber()
	result := n1 + n2
	res := &calculatorpb.SumResponse{
		SumResult:result,
	}
	return res, nil
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
