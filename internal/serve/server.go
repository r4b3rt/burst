package serve

import (
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/lxzan/gws"
	"google.golang.org/grpc"
	"log"
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
	connections    map[string]*connection
	wsPort         int64
}

func newServer(port int) *server {
	return &server{
		port:        port,
		connections: map[string]*connection{},
	}
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

	ws := gws.NewServer(s.e())
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}
	go func() {
		s.wsPort = int64(listener.Addr().(*net.TCPAddr).Port)
		err = ws.RunListener(listener)
		if err != nil {
			log.Fatal(err)
			return
		}
	}()

	return grpcServer.Serve(lis)
}

func (s *server) e() (gws.Event, *gws.ServerOption) {
	return &q{
		s: s,
	}, nil
}
