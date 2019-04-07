package main

import (
	"google.golang.org/grpc"
	"log"
	"grpc-go-course/calculator/calculatorpb"
	"context"
	"fmt"
	"io"
)

func main() {
	cc, err := grpc.Dial("localhost: 50052", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("err: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	//doUnary(c)
	doSeverStreaming(c)
}

func doSeverStreaming(c calculatorpb.CalculatorServiceClient){
	fmt.Println("start to do a Server Streaming RPC...")
	req := &calculatorpb.PrimeNumberDecompositionRequest{
		PrimeNumber: 120,
	}
	decompositionResStream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil{
		log.Fatalf("error while exec PrimeNumberDecomposition %v\n", err)
	}
	for{
		msg, err := decompositionResStream.Recv()
		if err == io.EOF{
			// we've reached the end of the stream
			break
		}
		if err != nil{
			log.Fatalf("error while reading stream: %v\n", err)
		}
		log.Printf("Response from Decomposition: %v\n", msg.GetResult())
	}
}

func doUnary(c calculatorpb.CalculatorServiceClient){
	fmt.Println("start to do a Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber: 10,
		SecondNumber: 20,
	}
	rep, err := c.Sum(context.Background(), req)
	if err != nil{
		log.Fatalf("err %v", err)
	}
	fmt.Println(rep.SumResult)
}