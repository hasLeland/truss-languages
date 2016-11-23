package gateway

import (
	"context"
	"fmt"

	svc "github.com/lelandbatey/truss-languages/gateway/gateway-service"

	swedish "github.com/lelandbatey/truss-languages/swedish/swedish-service"
	swedishhttp "github.com/lelandbatey/truss-languages/swedish/swedish-service/generated/client/http"

	canadian "github.com/lelandbatey/truss-languages/canadian/canadian-service"
	canadianhttp "github.com/lelandbatey/truss-languages/canadian/canadian-service/generated/client/http"
)

var router = map[string]func(string) string{
	"swedish":  Swedish,
	"canadian": Canadian,
}

// Gateway routes requests to their intended translation service
func Route(req svc.TranslateRequest) svc.TranslateResponse {
	rv := svc.TranslateResponse{}
	if len(req.Languages) < 1 {
		return svc.TranslateResponse{
			Value: "Not enough languages passed",
		}

	}
	rv.Value = req.Phrase
	for _, lang := range req.Languages {
		if translatefunc, ok := router[lang]; ok {
			rv.Value = translatefunc(rv.Value)
		} else {
			// Default to returning the phrase provided
			rv := svc.TranslateResponse{
				Value: fmt.Sprintf("nolangfound: %s\n%s", req.Languages, req.Phrase),
			}
			return rv
		}
	}
	return rv
}

func Swedish(in string) string {
	swedsvc, err := swedishhttp.New(":5051")
	if err != nil {
		panic(err)
	}
	swed, err := swedsvc.Translate(context.Background(), &swedish.TranslateRequest{Phrase: in})
	if err != nil {
		panic(err)
	}
	return swed.Value
}

func Canadian(in string) string {
	canadnsvc, err := canadianhttp.New(":5052")
	if err != nil {
		panic(err)
	}
	canadn, err := canadnsvc.Translate(context.Background(), &canadian.TranslateRequest{Phrase: in})
	if err != nil {
		panic(err)
	}
	return canadn.Value
}
