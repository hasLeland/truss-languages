package gateway

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	svc "github.com/hasLeland/truss-languages/gateway/gateway-service"

	swedish "github.com/hasLeland/truss-languages/swedish/swedish-service"
	swedishgrpc "github.com/hasLeland/truss-languages/swedish/swedish-service/generated/client/grpc"
	swedishhttp "github.com/hasLeland/truss-languages/swedish/swedish-service/generated/client/http"

	canadian "github.com/hasLeland/truss-languages/canadian/canadian-service"
	canadiangrpc "github.com/hasLeland/truss-languages/canadian/canadian-service/generated/client/grpc"
	canadianhttp "github.com/hasLeland/truss-languages/canadian/canadian-service/generated/client/http"
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
	httpsvc := func(port string) swedish.SwedishServer {
		fmt.Printf("Creating http connection for swedish!\n")
		service, err := swedishhttp.New(port)
		if err != nil {
			panic(err)
		}
		return service
	}
	grpcsvc := func(port string) swedish.SwedishServer {
		fmt.Printf("Creating grpc connection for swedish!\n")
		conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			panic(err)
		}
		service, err := swedishgrpc.New(conn)
		if err != nil {
			panic(err)
		}
		return service
	}
	_ = httpsvc
	_ = grpcsvc
	//swedsvc := httpsvc(":5051")
	swedsvc := grpcsvc(":5041")

	//swedsvc, err := swedishhttp.New(":5051")
	swed, err := swedsvc.Translate(context.Background(), &swedish.TranslateRequest{Phrase: in})
	if err != nil {
		panic(err)
	}
	return swed.Value
}

func Canadian(in string) string {
	httpsvc := func(port string) canadian.CanadianServer {
		fmt.Printf("Creating http connection for canadian!\n")
		service, err := canadianhttp.New(port)
		if err != nil {
			panic(err)
		}
		return service
	}
	grpcsvc := func(port string) canadian.CanadianServer {
		fmt.Printf("Creating grpc connection for canadian!\n")
		conn, err := grpc.Dial(port, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
		if err != nil {
			panic(err)
		}
		service, err := canadiangrpc.New(conn)
		if err != nil {
			panic(err)
		}
		return service
	}
	_ = httpsvc
	_ = grpcsvc
	canadnsvc := grpcsvc(":5042")
	//canadnsvc := httpsvc(":5052")

	//canadnsvc, err := canadianhttp.New(":5052")
	//conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	canadn, err := canadnsvc.Translate(context.Background(), &canadian.TranslateRequest{Phrase: in})
	if err != nil {
		panic(err)
	}
	return canadn.Value
}
