package main

import (
	"log"
	"net"

	protos "statusUrls/pkg/api/grpc/filegrps"

	"google.golang.org/grpc"
)

func main() {

	l, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal("%v", err)
	}

	cs := &protos.UnimplementedCheckServer{}

	//log := log.Default()
	// cs := server.NewGRPCServer(*log)
	//t := &server.GRPCServer{}

	grpc := grpc.NewServer()
	protos.RegisterCheckServer(grpc, cs)

	if err := grpc.Serve(l); err != nil {
		log.Fatal("%v", err)
	}

}
