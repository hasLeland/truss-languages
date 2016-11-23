// Package grpc provides a gRPC client for the Canadian service.
package grpc

import (
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/lelandbatey/truss-languages/canadian/canadian-service"
	svc "github.com/lelandbatey/truss-languages/canadian/canadian-service/generated"
)

// New returns an service backed by a gRPC client connection. It is the
func New(conn *grpc.ClientConn, options ...ClientOption) (pb.CanadianServer, error) {
	var cc clientConfig

	for _, f := range options {
		err := f(&cc)
		if err != nil {
			return nil, errors.Wrap(err, "cannot apply option")
		}
	}

	clientOptions := []grpctransport.ClientOption{
		grpctransport.ClientBefore(
			contextValuesToGRPCMetadata(cc.headers)),
	}
	var translateEndpoint endpoint.Endpoint
	{
		translateEndpoint = grpctransport.NewClient(
			conn,
			"canadian.Canadian",
			"Translate",
			svc.EncodeGRPCTranslateRequest,
			svc.DecodeGRPCTranslateResponse,
			pb.TranslateResponse{},
			clientOptions...,
		).Endpoint()
	}

	return svc.Endpoints{
		TranslateEndpoint: translateEndpoint,
	}, nil
}

type clientConfig struct {
	headers []string
}

// ClientOption is a function that modifies the client config
type ClientOption func(*clientConfig) error

func CtxValuesToSend(keys ...string) ClientOption {
	return func(o *clientConfig) error {
		o.headers = keys
		return nil
	}
}

func contextValuesToGRPCMetadata(keys []string) grpctransport.RequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		var pairs []string
		for _, k := range keys {
			if v, ok := ctx.Value(k).(string); ok {
				pairs = append(pairs, k, v)
			}
		}

		if pairs != nil {
			*md = metadata.Join(*md, metadata.Pairs(pairs...))
		}

		return ctx
	}
}
