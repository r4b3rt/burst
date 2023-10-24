package serve

import (
	"fmt"
	"github.com/fzdwx/burst/api"
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
		serverSideConn, err := net.Listen("tcp", fmt.Sprintf(":%d", m.ServerPort))
		if err != nil {
			eCh <- err
			return
		}

		doneCh <- s.react(peer, m, serverSideConn)
	}(peerTcpAddr, m)

	select {
	case id := <-doneCh:
		return id, nil
	case err := <-eCh:
		return "", err
	}
}
