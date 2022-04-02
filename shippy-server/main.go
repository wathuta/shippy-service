package main

import (
	"log"
	"net"

	pb "github.com/shippy/shippy-service/proto/consignment"
	"github.com/shippy/shippy-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

func main() {
	listener, err := net.Listen("tcp", port)
	checkError(err)

	s := grpc.NewServer()
	svr := service.NewShippingServiceServer()
	pb.RegisterShippingServiceServer(s, svr)

	reflection.Register(s)
	log.Printf("Listening on port %s", port)
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
