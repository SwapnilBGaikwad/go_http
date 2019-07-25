package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "projects/go_http/product/product"
)

const (
	address     = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPriceCalculatorClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	products := []*pb.Product{
		{Name: "Product 1", Price: 10},
		{Name: "Product 2", Price: 40},
	}
	r, err := c.CalculateTotal(ctx, &pb.ProductRequest{Products: products})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Total of Products is: %f", r.Total)
}
