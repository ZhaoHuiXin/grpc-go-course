package main

import (
	"fmt"
	"log"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"context"
	"io"
)

func main()  {
	fmt.Println("hello i'a a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	//fmt.Printf("Created client: %f", c)

	//doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Max",
			LastName: "zhao",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil{
		log.Fatalf("error while calliong Greeting RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do Server Streaming RPC...")

	req := &greetpb.GreetManyTimesRequest{
		Greeting:&greetpb.Greeting{
			FirstName: "haha",
			LastName: "heihei",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil{
		log.Fatalf("error while calling Server Streaming Greeting RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF{
			// we've reached the end of the stream
			break
		}
		if err != nil{
			log.Fatalf("error while reading stream: %v\n", err)
		}
		log.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}

}
