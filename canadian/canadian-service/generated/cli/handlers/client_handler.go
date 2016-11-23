package clienthandler

import (
	pb "github.com/lelandbatey/truss-languages/canadian/canadian-service"
)

// Translate implements Service.
func Translate(PhraseTranslate string) (*pb.TranslateRequest, error) {
	request := pb.TranslateRequest{
		Phrase: PhraseTranslate,
	}
	return &request, nil
}
