package serve

import (
	"fmt"
	"github.com/fzdwx/burst/api"
	"net"
)

func (s *server) mapping(peer net.Addr, m *api.PortMapping) error {
	eCh := make(chan error)
	doneCh := make(chan struct{})

	peerTcpAddr, ok := peer.(*net.TCPAddr)
	if !ok {
		return fmt.Errorf("peer not tcp addr")
	}

	go func(peer *net.TCPAddr, m *api.PortMapping) {
		serverSideConn, err := net.Listen("tcp", fmt.Sprintf(":%d", m.ServerPort))
		if err != nil {
			eCh <- err
			return
		}

		s.react(peer, m, serverSideConn)

		doneCh <- struct{}{}
	}(peerTcpAddr, m)

	select {
	case <-doneCh:
		return nil
	case err := <-eCh:
		return err
	}
}
