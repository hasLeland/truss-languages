package handler

import (
	"strings"

	"golang.org/x/net/context"

	pb "github.com/lelandbatey/truss-languages/swedish/swedish-service"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.SwedishServer {
	return swedishService{}
}

type swedishService struct{}

// Translate implements Service.
func (s swedishService) Translate(ctx context.Context, in *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	var resp pb.TranslateResponse
	// A nice trick to translate most languages to Swedish
	resp.Value = strings.Join(strings.Split(in.Phrase, ""), "f")
	//for _, letter := range in.Phrase {
	//resp.Value += string(letter) + "f"
	//}
	return &resp, nil
}
