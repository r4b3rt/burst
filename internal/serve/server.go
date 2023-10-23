package serve

import (
	context "context"
	"fmt"
	"github.com/fzdwx/burst/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log/slog"
	"net"
)

func ListenAndServe(port int) error {
	s := newServer(port)
	return s.ListenAndServe()
}

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

	slog.Info("start burst server", slog.Int("port", s.port))
	return grpcServer.Serve(lis)
}

func (s *server) Export(ctx context.Context, request *api.ExportRequest) (*api.ExportResponse, error) {
	p, ok := peer.FromContext(ctx)
	if ok == false {
		return &api.ExportResponse{}, fmt.Errorf("peer not found")
	}
	slog.Info("handle export", slog.String("portMapping", toString(request.PortMapping)), slog.String("clientAddr", p.Addr.String()))

	return &api.ExportResponse{}, nil
}

func toString(mapping []*api.PortMapping) string {
	var str string
	for i, m := range mapping {
		if i != 0 {
			str += ", "
		}
		str += fmt.Sprintf("%d to %d", m.ClientPort, m.ServerPort)
	}
	return str
}
