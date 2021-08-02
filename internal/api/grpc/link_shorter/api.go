package linkshorter

import (
	"context"
	"fmt"
	"ozon/internal/api/grpc/gen"
	"ozon/internal/service"
)

// GRPCServer implements gRPC generated
// interface gen.LinkShorterServer
type GRPCServer struct {
	gen.UnimplementedLinkShorterServer
}

// Create - method for creating short link.
// Incoming args:
//	- LongLink: it will be short cut.
// Outcoming params:
//	- ShortLink: modified short link.
func (s *GRPCServer) Create(ctx context.Context, in *gen.CreateRequest) (*gen.CreateResponse, error) {
	inLink := in.GetLongLink()

	if inLink == "" {
		return nil, fmt.Errorf("incoming link is empty")
	}

	shortLink, err := service.Create(ctx, inLink)
	if err != nil {
		return nil, err
	}

	return &gen.CreateResponse{
		ShortLink: shortLink,
	}, nil
}

// Get - method for getting original link through short one.
// Incoming args:
//	- ShortLink: will be used to get original link.
// Outcoming params:
//	- LongLink: original link.
func (s *GRPCServer) Get(ctx context.Context, in *gen.GetRequest) (*gen.GetResponse, error) {
	inLink := in.GetShortLink()

	if inLink == "" {
		return nil, fmt.Errorf("incoming link is empty")
	}

	longLink, err := service.Get(ctx, inLink)
	if err != nil {
		return nil, err
	}

	return &gen.GetResponse{
		LongLink: longLink,
	}, nil
}
