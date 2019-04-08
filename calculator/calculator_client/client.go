package main

import (
	"google.golang.org/grpc"
	"log"
	"grpc-go-course/calculator/calculatorpb"
	"context"
	"fmt"
	"io"
	"time"
)

func main() {
	cc, err := grpc.Dial("localhost: 50052", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("err: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	//doUnary(c)
	//doSeverStreaming(c)
	//doClientStreaming(c)
	doBiDiStreaming(c)
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

func doClientStreaming(c calculatorpb.CalculatorServiceClient){
	stream, err := c.ComputeAverage(context.Background())
	if err != nil{
		log.Fatalf("error while ComputeAverage %v\n", err)
	}
	requests := []*calculatorpb.ComputeAverageRequest{
		&calculatorpb.ComputeAverageRequest{
			InputNumber: int32(200),
		},
		&calculatorpb.ComputeAverageRequest{
			InputNumber: int32(4),
		},
		&calculatorpb.ComputeAverageRequest{
			InputNumber: int32(4),
		},
		&calculatorpb.ComputeAverageRequest{
			InputNumber: int32(4),
		},
	}
	for _, req := range requests{
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("error while CloseAndRecv %v\n", err)
	}else{
		log.Printf("ComputeAverage Response: %v\n", res.GetAverageNumber())
	}
}

func doBiDiStreaming(c calculatorpb.CalculatorServiceClient){
	stream, err := c.FindMaximum(context.Background())
	if err != nil{
		log.Fatalf("error while create FindMaxNumber Client %v\n", err)
		return
	}
	dataSend := []int32{1,5,3,6,2,20}
	waitc := make(chan struct{}) // wait channel

	// function to send a bunch of messages
	go func() {
		for _, n := range dataSend{
			stream.Send(&calculatorpb.FindMaximumRequest{
				InputNumber: n,
			})
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// function to receive a bunch of messages
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF{
				break
			}
			if err != nil{
				log.Fatalf("error while client stream Recv %v\n", err)
				break
			}
			fmt.Printf("now the MaxNumber is: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	<- waitc
}