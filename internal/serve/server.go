package serve

import (
	context "context"
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type server struct {
	api.BurstServer

	port int
}

func newServer(port int) *server {
	return &server{port: port}
}

func (s *server) ListenAndServe() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.port))
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterBurstServer(grpcServer, s)

	slog.Info("Start burst", slog.Int("port", s.port))
	return grpcServer.Serve(lis)
}

func (s *server) Export(ctx context.Context, request *api.ExportRequest) (*api.ExportResponse, error) {
	return nil, nil
}

func ListenAndServe(cmd *cobra.Command, args []string) error {
	s := newServer(8000)
	return s.ListenAndServe()
}
