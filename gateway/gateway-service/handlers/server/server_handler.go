package handler

import (
	"golang.org/x/net/context"

	pb "github.com/hasLeland/truss-languages/gateway/gateway-service"

	impl "github.com/hasLeland/truss-languages/gateway/gateway-impl"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.GatewayServer {
	return gatewayService{}
}

type gatewayService struct{}

// Translate implements Service.
func (s gatewayService) Translate(ctx context.Context, in *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	var resp pb.TranslateResponse
	resp = impl.Route(*in)
	return &resp, nil
}
