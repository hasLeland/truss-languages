package svc

// This file provides server-side bindings for the gRPC transport.
// It utilizes the transport/grpc.Server.

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	grpctransport "github.com/go-kit/kit/transport/grpc"

	// This Service
	pb "github.com/hasLeland/truss-languages/swedish/swedish-service"
)

// MakeGRPCServer makes a set of endpoints available as a gRPC SwedishServer.
func MakeGRPCServer(ctx context.Context, endpoints Endpoints) pb.SwedishServer {
	serverOptions := []grpctransport.ServerOption{
		grpctransport.ServerBefore(metadataToContext),
	}
	return &grpcServer{
		// swedish

		translate: grpctransport.NewServer(
			ctx,
			endpoints.TranslateEndpoint,
			DecodeGRPCTranslateRequest,
			EncodeGRPCTranslateResponse,
			serverOptions...,
		),
	}
}

// grpcServer implements the SwedishServer interface
type grpcServer struct {
	translate grpctransport.Handler
}

// Methods for grpcServer to implement SwedishServer interface

func (s *grpcServer) Translate(ctx context.Context, req *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	_, rep, err := s.translate.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.TranslateResponse), nil
}

// Server Decode

// DecodeGRPCTranslateRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC translate request to a user-domain translate request. Primarily useful in a server.
func DecodeGRPCTranslateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.TranslateRequest)
	return req, nil
}

// Client Decode

// DecodeGRPCTranslateResponse is a transport/grpc.DecodeResponseFunc that converts a
// gRPC translate reply to a user-domain translate response. Primarily useful in a client.
func DecodeGRPCTranslateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*pb.TranslateResponse)
	return reply, nil
}

// Server Encode

// EncodeGRPCTranslateResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain translate response to a gRPC translate reply. Primarily useful in a server.
func EncodeGRPCTranslateResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.TranslateResponse)
	return resp, nil
}

// Client Encode

// EncodeGRPCTranslateRequest is a transport/grpc.EncodeRequestFunc that converts a
// user-domain translate request to a gRPC translate request. Primarily useful in a client.
func EncodeGRPCTranslateRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.TranslateRequest)
	return req, nil
}

// Helpers

func metadataToContext(ctx context.Context, md *metadata.MD) context.Context {
	for k, v := range *md {
		if v != nil {
			// The key is added both in metadata format (k) which is all lower
			// and the http.CanonicalHeaderKey of the key so that it can be
			// accessed in either format
			ctx = context.WithValue(ctx, k, v[0])
			ctx = context.WithValue(ctx, http.CanonicalHeaderKey(k), v[0])
		}
	}

	return ctx
}
