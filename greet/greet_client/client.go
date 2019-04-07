package main

import (
	"fmt"
	"log"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"context"
	"io"
	"time"
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
	//doServerStreaming(c)
	doClientStreaming(c)
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

func doClientStreaming(c greetpb.GreetServiceClient){
	fmt.Println("Starting to do Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "AAA",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "BBB",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "CCC",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "DDD",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "zhao",
			},
		},
	}
	stream , err := c.LongGreet(context.Background() )
	if err != nil{
		log.Fatalf("error while calling LongGreet: %v", err)
	}
	// we iterate over our slice and send each message individually
	for _, req := range requests{
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(3 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil{
		log.Fatalf("error while receiving response from LongGreet: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}
