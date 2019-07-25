//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/product.proto

package main

import (
	"golang.org/x/net/context"
	pb "projects/go_http/product/product"
	"testing"
)

type test struct {
	server *server
}

func Test_server_CalculateTotal(t *testing.T) {
	server := test{
		server: &server{},
	}
	s := server.server

	request := &pb.ProductRequest{Products: []*pb.Product{
		{Name: "P1", Price: 1.0},
		{Name: "P2", Price: 4.0},
	}}

	got, _ := s.CalculateTotal(context.Background(), request)

	if got.Total != 5.0 {
		t.Errorf("server.CalculateTotal() = %v, want %v", got, 0)
	}
}
