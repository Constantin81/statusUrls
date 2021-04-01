package server

import (
	"context"
	"fmt"
	"log"
	protos "statusUrls/pkg/api/grpc/filegrps"
)

//...
type GRPCServer struct {
	log.Logger
}

func NewGRPCServer(l log.Logger) *GRPCServer {
	return &GRPCServer{l}
}

// ...
func (s *GRPCServer) AddUrl(ctx context.Context, req *protos.AddRequestUrl) (*protos.AddResponseUrl, error) {
	sampleString := fmt.Sprint(req.CountPointCheckUrl)
	return &protos.AddResponseUrl{Result: req.StrUrl + sampleString}, nil
}

func (s *GRPCServer) GetStatusUrl(ctx context.Context, req *protos.RequestInfoByCheckUrl) (*protos.ResponseInfoByCheckUrl, error) {
	return &protos.ResponseInfoByCheckUrl{TimeCheckUrl: req.StrUrl, StatusUrl: req.StrUrl}, nil
}

func (s *GRPCServer) DeleteUrl(ctx context.Context, req *protos.RequestUrlDelete) (*protos.ResponseUrlDelete, error) {
	return &protos.ResponseUrlDelete{StrUrl: req.StrUrl}, nil
}

func (s *GRPCServer) mustEmbedUnimplementedCheckServer() {
}
