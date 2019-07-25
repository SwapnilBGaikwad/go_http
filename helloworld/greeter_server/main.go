//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"math"
	"net"
	pb "projects/go_http/helloworld/helloworld"
	"strconv"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{
	Count int
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server)SayHelloAgain(ctx context.Context,in *pb.HelloRequest) (*pb.HelloReply, error) {
	s.Count++
	pow := math.Pow(2, float64(s.Count))
	return &pb.HelloReply{Message: "Pow : " + strconv.FormatFloat(pow, 'f', 6, 64)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Started server on port " + port)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
