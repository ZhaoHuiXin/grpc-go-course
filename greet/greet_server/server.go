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