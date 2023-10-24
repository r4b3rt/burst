package serve

import (
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/util/id"
	"net"
)

type connection struct {
	peer net.Addr
	id   string

	serverSideConn net.Listener
	mapping        *api.PortMapping
}

func (s *server) react(peer net.Addr, m *api.PortMapping, serverSideConn net.Listener) string {
	c := s.addConnection(peer, m, serverSideConn)

	switch serverSideConn.(type) {
	case *net.TCPListener:
		s.reactTCP(peer, m, serverSideConn.(*net.TCPListener))
	}

	return c.id
}

func (s *server) reactTCP(peer net.Addr, m *api.PortMapping, conn *net.TCPListener) {
	m.ServerPort = int64(conn.Addr().(*net.TCPAddr).Port)

	// todo react
}

func (s *server) addConnection(peer net.Addr, m *api.PortMapping, serverSideConn net.Listener) *connection {
	s.connectionLock.Lock()
	defer s.connectionLock.Unlock()

	c := &connection{
		peer:           peer,
		id:             id.Next(),
		serverSideConn: serverSideConn,
		mapping:        m,
	}
	s.connections[c.id] = c
	return c
}

func (s *server) getConnection(id string) *connection {
	s.connectionLock.RLock()
	defer s.connectionLock.RUnlock()

	return s.connections[id]
}
