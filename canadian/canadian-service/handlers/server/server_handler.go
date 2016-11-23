package handler

import (
	"strings"

	"golang.org/x/net/context"

	pb "github.com/lelandbatey/truss-languages/canadian/canadian-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.CanadianServer {
	return canadianService{}
}

type canadianService struct{}

// Translate implements Service.
func (s canadianService) Translate(ctx context.Context, in *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	var resp pb.TranslateResponse
	// A neet trick to translate to canadian
	resp.Value = strings.Join(strings.Split(in.Phrase, " "), "ay ")
	return &resp, nil
}