package main

import (
	"fmt"
	"log"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
)

func main()  {
	fmt.Println("hello i'a a client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)

}
