package serve

import (
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/internal/cons"
	"github.com/fzdwx/burst/util/log"
	"log/slog"
	"net"
)

func (s *server) mapping(peer net.Addr, m *api.PortMapping) (string, error) {
	eCh := make(chan error)
	doneCh := make(chan string)

	peerTcpAddr, ok := peer.(*net.TCPAddr)
	if !ok {
		return "", fmt.Errorf("peer not tcp addr")
	}

	go func(peer *net.TCPAddr, m *api.PortMapping) {
		switch m.Protocol {
		case cons.Protocol.TCP:
			serverSideConn, err := net.Listen("tcp", fmt.Sprintf(":%d", m.ServerPort))
			if err != nil {
				eCh <- err
				return
			}
			c := s.addConnection(peer, m, serverSideConn)
			m.ServerPort = int64(serverSideConn.(*net.TCPListener).Addr().(*net.TCPAddr).Port)

			slog.Info("success listen tcp", log.ConnectionId(c.id), log.Mapping(c.mapping))
			doneCh <- c.id
		default:
			eCh <- fmt.Errorf("unsupported protocol")
		}
	}(peerTcpAddr, m)

	select {
	case id := <-doneCh:
		return id, nil
	case err := <-eCh:
		return "", err
	}
}
