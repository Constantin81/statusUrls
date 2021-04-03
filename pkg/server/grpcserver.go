package server

import (
	"context"
	"errors"
	"fmt"
	"log"
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
	TimeCheckUrl string
}

var List = []InfoAboutUrl{}

// ...
func (s *GRPCServer) AddUrl(ctx context.Context, req *protos.AddRequestUrl) (*protos.AddResponseUrl, error) {

	_, err := url.ParseRequestURI(req.StrUrl)

	if err != nil {
		//return &protos.AddResponseUrl{Result: "No valid url"}, nil
		log.Fatal(err)
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

	//strLastInfoUrl := make([]TimeCheckAndStatusUrl, 5)
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
			return &protos.ResponseInfoByCheckUrl{}, err
		}
	}
	return &protos.ResponseInfoByCheckUrl{Containers: containerInfo}, nil
}

func (s *GRPCServer) DeleteUrl(ctx context.Context, req *protos.RequestUrlDelete) (*protos.ResponseUrlDelete, error) {
	return &protos.ResponseUrlDelete{StrUrl: req.StrUrl}, nil
}

func (s *GRPCServer) mustEmbedUnimplementedCheckServer() {
}

// func working(data []*protos.Container) ([]*protos.Container, error) {
// 	var cs []*protos.Container
// 	err := json.Unmarshal([]byte(data), &cs)
// 	if err != nil {
// 		log.Fatal(new.error("Can`t decode data"))
// 	}
// 	// handle the error here
// 	return cs, nil
// }

//*resp, err := http.Get(req.StrUrl)
// if err != nil {
// 	//log.Fatal(err)
// 	return &protos.ResponseInfoByCheckUrl{TimeCheckUrl: req.StrUrl, StatusUrl: req.StrUrl}, nil
// }

// // Print the HTTP Status Code and Status Name
// fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
// fmt.Println("HTTP Response Status:", )

// if resp.StatusCode >= 200 && resp.StatusCode <= 299 {

// 	fmt.Println("HTTP Status is in the 2xx range")
// } else {
// 	fmt.Println("Argh! Broken")
//}
//
