package linkshorter

import (
	"context"
	"ozon/internal/api/grpc/gen"
	"ozon/internal/service"
)

// GRPCServer implements gRPC generated
// interface gen.LinkShorterServer
type GRPCServer struct {
	gen.UnimplementedLinkShorterServer
}

func (s *GRPCServer) Create(ctx context.Context, in *gen.CreateRequest) (*gen.CreateResponse, error) {
	shortLink, err := service.Create(ctx, "")
	if err != nil {
		return nil, err
	}

	return &gen.CreateResponse{
		ShortLink: shortLink,
	}, nil
}

func (s *GRPCServer) Get(ctx context.Context, in *gen.GetRequest) (*gen.GetResponse, error) {
	longLink, err := service.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	return &gen.GetResponse{
		LongLink: longLink,
	}, nil
}
