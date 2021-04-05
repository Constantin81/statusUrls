package main

import (
	"log"
	"net"

	pb "statusUrls/api/currency"
	protos "statusUrls/pkg/server"

	"google.golang.org/grpc"
)

func main() {

	var site = new(protos.InfoAboutUrl)
	site.SiteUrl = "http://google.com"
	site.CountCheck = 1

	status := []protos.TimeCheckAndStatusUrl{}
	for i := 0; i < 7; i++ {
		n := protos.TimeCheckAndStatusUrl{StatusCode: int32(200), TimeCheckUrl: "2021-01-31 00:00:00 +0000 UTC"}
		status = append(status, n)
	}

	site.InfoUrlAboutCheck = status

	protos.List = append(protos.List, *site)

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	cs := &protos.GRPCServer{}

	grpc := grpc.NewServer()
	pb.RegisterCheckServer(grpc, cs)

	if err := grpc.Serve(l); err != nil {
		log.Fatal(err)
	}

}
