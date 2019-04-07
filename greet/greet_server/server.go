package main
import (
	"context"
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
	"strconv"
	"time"
	"io"
)

type server struct{}
// implement func Greet
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error){
	fmt.Printf("Greet function was invoked with %v \n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest,
	stream greetpb.GreetService_GreetManyTimesServer) error{
		fmt.Printf("GreetManyTimes func wai invoked with %v\n", req)
		firstName := req.GetGreeting().GetFirstName()
		for i := 0; i < 10; i++ {
			result := "hello " + firstName + " number" + strconv.Itoa(i)
			res := &greetpb.GreetManyTimesResponse{
				Result: result,
			}
			stream.Send(res)
			time.Sleep(3 * time.Second)
	}
	return nil
}

func (*server)LongGreet(stream greetpb.GreetService_LongGreetServer) error{
	fmt.Println("LongGreet func was invoked with a streaminbg request")
	result := ""
	for{
		req, err := stream.Recv()
		if err == io.EOF{
			// we have finished reading the client stream
			//break
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil{
			log.Fatalf("error while reading client stream: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result +="hello " + firstName + "!"
	}
}

func (*server)GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error{
	stream.Send(&greetpb.GreetEveryoneResponse{
		Result: "i got you before you call me.",
	})
	for {
		req, err := stream.Recv()
		if err == io.EOF{
			break
		}else if err != nil{
			log.Fatalf("error when Recv %v\n", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		if firstName == "lucy"{
			result := "Welcome back " + firstName + ", my lord! "
			stream.Send(&greetpb.GreetEveryoneResponse{
				Result: result,
			})
		}else{
			result := "Hello " + firstName + "! "
			stream.Send(&greetpb.GreetEveryoneResponse{
				Result: result,
			})
		}
	}
	return nil
}

func main(){
	fmt.Println("hello world")

	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}