package server

import (
	"HTTP-SERVER/api"
	"context"
	"statusUrls/pkg/api"
)

//...
type GRPCServer struct {
}

// ...
func (s *GRPCServer) AddUrl(ctx context.Context, req *api.AddRequestUrl) (*api.AddResponseUrl, error) {
	return &api.AddResponseUrl{Result: req.strUrl}, nil
}
