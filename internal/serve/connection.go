package serve

import (
	"github.com/fzdwx/burst/api"
	"net"
)

type connection struct {
	peer net.Addr

	serverConn net.Listener
	mapping    *api.PortMapping
}

func (s *server) react(peer net.Addr, m *api.PortMapping, conn net.Listener) {
	s.connectionLock.Lock()
	s.connections = append(s.connections, &connection{
		peer:       peer,
		serverConn: conn,
		mapping:    m,
	})
	s.connectionLock.Unlock()

	switch conn.(type) {
	case *net.TCPListener:
		s.reactTCP(peer, m, conn.(*net.TCPListener))
	}
}

func (s *server) reactTCP(peer net.Addr, m *api.PortMapping, conn *net.TCPListener) {
	m.ServerPort = int64(conn.Addr().(*net.TCPAddr).Port)

	// todo react
}
