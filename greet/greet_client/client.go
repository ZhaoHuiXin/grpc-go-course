package main

import (
	"fmt"
	"log"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"context"
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

	doUnary(c)
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
