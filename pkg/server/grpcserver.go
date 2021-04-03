package server

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	protos "statusUrls/api/currency"
)

//...
type GRPCServer struct {
	protos.UnimplementedCheckServer
}

type InfoAboutUrl struct {
	SiteUrl           string
	CountCheck        int32
	InfoUrlAboutCheck []TimeCheckAndStatusUrl
}

type TimeCheckAndStatusUrl struct {
	StatusCode   int32
	TimeCheckUrl int32
}

var List = []InfoAboutUrl{}

// ...
func (s *GRPCServer) AddUrl(ctx context.Context, req *protos.AddRequestUrl) (*protos.AddResponseUrl, error) {

	_, err := url.ParseRequestURI(req.StrUrl)

	if err != nil {
		return &protos.AddResponseUrl{Result: "No valid url"}, nil
		//log.Fatal(err)
	}

	var site = new(InfoAboutUrl)
	site.SiteUrl = req.StrUrl
	if req.CountPointCheckUrl == 0 {
		site.CountCheck = 1
	} else {
		site.CountCheck = req.CountPointCheckUrl
	}

	List = append(List, *site)
	fmt.Println("List Urls = ", List)

	return &protos.AddResponseUrl{Result: fmt.Sprint("Url = ", req.StrUrl, " added")}, nil
}

func (s *GRPCServer) GetStatusUrl(ctx context.Context, req *protos.RequestInfoByCheckUrl) (*protos.ResponseInfoByCheckUrl, error) {
	resp, err := http.Get(req.StrUrl)
	if err != nil {
		//log.Fatal(err)
		return &protos.ResponseInfoByCheckUrl{TimeCheckUrl: req.StrUrl, StatusUrl: req.StrUrl}, nil
	}

	// Print the HTTP Status Code and Status Name
	fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {

		fmt.Println("HTTP Status is in the 2xx range")
	} else {
		fmt.Println("Argh! Broken")
	}

	return &protos.ResponseInfoByCheckUrl{TimeCheckUrl: req.StrUrl, StatusUrl: req.StrUrl}, nil
}

func (s *GRPCServer) DeleteUrl(ctx context.Context, req *protos.RequestUrlDelete) (*protos.ResponseUrlDelete, error) {
	return &protos.ResponseUrlDelete{StrUrl: req.StrUrl}, nil
}

func (s *GRPCServer) mustEmbedUnimplementedCheckServer() {
}
