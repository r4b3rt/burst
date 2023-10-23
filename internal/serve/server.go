package serve

import (
	"fmt"
	"github.com/fzdwx/burst/api"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"sync"
)

func ListenAndServe(port int) error {
	s := newServer(port)
	return s.ListenAndServe()
}

type server struct {
	api.BurstServer

	port           int
	connectionLock sync.RWMutex
	connections    []*connection
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
