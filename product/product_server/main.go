//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/product.proto

package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "projects/go_http/product/product"
)

const (
	port = ":50051"
)

// server is used to implement product.PriceCalculatorServer.
type server struct{
}

func (s *server)CalculateTotal(ctx context.Context, in *pb.ProductRequest) (*pb.ProductReply, error)  {
	var total float32
	for _, product := range in.Products {
		total += product.Price
	}
	return &pb.ProductReply{
		Total: total,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Started server on port " + port)
	}
	s := grpc.NewServer()
	pb.RegisterPriceCalculatorServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
