package main

import (
	"log"
	"net"

	pb "statusUrls/api/currency"
	protos "statusUrls/pkg/server"

	"google.golang.org/grpc"
)

func main() {

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("%v", err)
	}

	cs := &protos.GRPCServer{}

	grpc := grpc.NewServer()
	pb.RegisterCheckServer(grpc, cs)

	if err := grpc.Serve(l); err != nil {
		log.Fatal("%v", err)
	}

}
