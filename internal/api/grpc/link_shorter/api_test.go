package linkshorter_test

import (
	"context"
	"log"
	"net"
	"ozon/internal/api/grpc/gen"
	linkshorter "ozon/internal/api/grpc/link_shorter"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var (
	lis        *bufconn.Listener
	respCreate *gen.CreateResponse
)

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	gen.RegisterLinkShorterServer(s, &linkshorter.GRPCServer{})

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := gen.NewLinkShorterClient(conn)
	resp, err := client.Create(ctx, &gen.CreateRequest{LongLink: "link"})
	if err != nil {
		t.Fatalf("CreateRequest failed: %v", err)
	}

	respCreate = resp

	log.Printf("Response: %+v", resp)
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := gen.NewLinkShorterClient(conn)
	resp, err := client.Get(ctx, &gen.GetRequest{ShortLink: respCreate.GetShortLink()})
	if err != nil {
		t.Fatalf("GetRequest failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
}
