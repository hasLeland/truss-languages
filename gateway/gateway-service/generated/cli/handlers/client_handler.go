package clienthandler

import (
	pb "github.com/hasLeland/truss-languages/gateway/gateway-service"
)

// Translate implements Service.
func Translate(PhraseTranslate string, LanguagesTranslate []string) (*pb.TranslateRequest, error) {
	request := pb.TranslateRequest{
		Phrase:    PhraseTranslate,
		Languages: LanguagesTranslate,
	}
	return &request, nil
}
