package main

import (
	"google.golang.org/grpc"
	"github.com/prometheus/common/log"
	"grpc-go-course/calculator/calculatorpb"
	"context"
	"fmt"
)

func main() {
	cc, err := grpc.Dial("localhost: 50052", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("err: %v", err)
	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient){
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