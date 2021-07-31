package linkshorter

import (
	"context"
	"ozon/internal/api/grpc/gen"
)

// GRPCServer implements gRPC generated
// interface gen.LinkShorterServer
type GRPCServer struct {
	gen.UnimplementedLinkShorterServer
}

func (s *GRPCServer) Create(ctx context.Context, in *gen.CreateRequest) (*gen.CreateResponse, error) {
	return &gen.CreateResponse{}, nil
}

func (s *GRPCServer) Get(ctx context.Context, in *gen.GetRequest) (*gen.GetResponse, error) {
	return &gen.GetResponse{}, nil
}
