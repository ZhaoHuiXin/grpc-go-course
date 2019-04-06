package main
import (
	"fmt"
	"net"
	"log"
	"google.golang.org/grpc"
	"grpc-go-course/greet/greetpb"
)

type server struct{}

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