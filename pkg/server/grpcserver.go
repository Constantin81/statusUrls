package server

import (
	"context"
	"errors"
	"fmt"
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
	SearchableUrl     bool
	InfoUrlAboutCheck []TimeCheckAndStatusUrl
}

type TimeCheckAndStatusUrl struct {
	StatusCode   int32
	TimeCheckUrl string
}

var List = []InfoAboutUrl{}

// ...
func (s *GRPCServer) AddUrl(ctx context.Context, req *protos.AddRequestUrl) (*protos.AddResponseUrl, error) {

	_, err := url.ParseRequestURI(req.StrUrl)

	if err != nil {
		err := errors.New("No valid url")
		return &protos.AddResponseUrl{Result: ""}, err
		//log.Fatal(err)
	}

	var site = new(InfoAboutUrl)
	site.SiteUrl = req.StrUrl
	site.SearchableUrl = true
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

	containerInfo := []*protos.Container{}

	for _, s := range List {
		if s.SiteUrl == req.StrUrl {
			fmt.Println("Url = ", s.SiteUrl)

			for _, i := range s.InfoUrlAboutCheck {

				var infoUrl = &protos.Container{TimeCheckUrl: i.TimeCheckUrl, StatusUrl: int32(i.StatusCode)}

				containerInfo = append(containerInfo, infoUrl)

			}
			return &protos.ResponseInfoByCheckUrl{Containers: containerInfo}, nil
		} else {
			err := errors.New("math: square root of negative number")
			return &protos.ResponseInfoByCheckUrl{Containers: nil}, err
		}
	}
	return &protos.ResponseInfoByCheckUrl{Containers: containerInfo}, nil
}

func (s *GRPCServer) DeleteUrl(ctx context.Context, req *protos.RequestUrlDelete) (*protos.ResponseUrlDelete, error) {

	_, err := url.ParseRequestURI(req.StrUrl)

	if err != nil {
		err := errors.New("No valid url")
		return &protos.ResponseUrlDelete{StrUrl: ""}, err
	}

	for _, s := range List {
		if s.SiteUrl == req.StrUrl {
			s.SearchableUrl = false
			return &protos.ResponseUrlDelete{StrUrl: "Url - " + req.StrUrl + " removed from check"}, nil
		}
	}

	return &protos.ResponseUrlDelete{StrUrl: req.StrUrl}, nil
}

func (s *GRPCServer) mustEmbedUnimplementedCheckServer() {
}
