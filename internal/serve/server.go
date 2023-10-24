package serve

import (
	"bufio"
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/util/log"
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
	connections    map[string]*connection
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
	return grpcServer.Serve(lis)
}

func (s *server) transformUserToClient(userConn *userConnection, clientStream api.Burst_TransformServer) {
	r := bufio.NewReader(userConn.conn)
	buf := make([]byte, 1024)
	for {
		// 1. read user conn data
		n, err := r.Read(buf)
		if err != nil {
			slog.Error("read user conn error, stop read",
				log.UserToClient(),
				log.ConnectionId(userConn.clientConnectionId),
				log.UserConnectionId(userConn.id),
				log.Reason(err))
			return
		}

		// 2. send data to client stream
		if err = clientStream.Send(&api.TransFromData{
			ConnectionId:     userConn.clientConnectionId,
			UserConnectionId: userConn.id,
			Data:             buf[:n],
		}); err != nil {
			slog.Error("send data to client error, stop send",
				log.UserToClient(),
				log.ConnectionId(userConn.clientConnectionId),
				log.UserConnectionId(userConn.id),
				log.Reason(err))
			return
		}
	}

}
