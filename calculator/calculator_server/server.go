package main

import (
	"net"
	"log"
	"grpc-go-course/calculator/calculatorpb"
	"context"
	"fmt"
	"time"
	"io"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc"
	"math"
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

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error{
	var count, sum int32
	for {
		req, err := stream.Recv()

		if err == io.EOF{
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				AverageNumber: sum / count,
			})
		}
		if err != nil{
			log.Fatalf("error while stream.Recv %v\n", err)
		}
		sum += req.GetInputNumber()
		count += 1
	}
}

func (*server) FindMaximum( stream calculatorpb.CalculatorService_FindMaximumServer) error{
	maxNum := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatalf("error while FindMaximum server Recv %v\n", err)
		}
		log.Printf("Server receive number: %v\n", req.GetInputNumber())
		if  req.GetInputNumber() > maxNum{
			maxNum = req.GetInputNumber()
			stream.Send(&calculatorpb.FindMaximumResponse{
				Result: maxNum,
			})
		}
	}
	return nil
}


func (*server)SquareRoot(ctx context.Context, req *calculatorpb.SquareRootRequest) (*calculatorpb.SquareRootResponse, error){
	fmt.Println("Received SquareRoot RPC")
	number := req.GetNumber()
	if (number < 0){
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %v", number),
		)
	}
	return &calculatorpb.SquareRootResponse{
		NumberRoot: math.Sqrt(float64(number)),
	}, nil

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
