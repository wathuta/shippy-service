package service

import (
	"context"

	pb "github.com/shippy/shippy-service/proto/consignment"
)

type ShippingServiceServer interface {
	CreateConsignment(context.Context, *pb.Consignment) (*pb.Response, error)
	GetConsignments(context.Context, *pb.GetRequest) (*pb.Response, error)
	mustEmbedUnimplementedShippingServiceServer()
}

type Shipping_serviceServer struct {
	pb.UnimplementedShippingServiceServer
}

func NewShippingServiceServer() *Shipping_serviceServer {
	return &Shipping_serviceServer{}
}

func (s *Shipping_serviceServer) CreateConsignment(ctx context.Context, consignment *pb.Consignment) (*pb.Response, error) {
	var ret *pb.Response
	return ret, nil
}
func (s *Shipping_serviceServer) GetConsignments(ctx context.Context, consignment *pb.GetRequest) (*pb.Response, error) {
	var ret *pb.Response
	return ret, nil
}
