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

func (s *server) react(peer net.Addr, m *api.PortMapping, serverSideConn net.Listener) {
	s.connectionLock.Lock()
	s.connections = append(s.connections, &connection{
		peer:           peer,
		id:             id.Next(),
		serverSideConn: serverSideConn,
		mapping:        m,
	})
	s.connectionLock.Unlock()

	switch serverSideConn.(type) {
	case *net.TCPListener:
		s.reactTCP(peer, m, serverSideConn.(*net.TCPListener))
	}
}

func (s *server) reactTCP(peer net.Addr, m *api.PortMapping, conn *net.TCPListener) {
	m.ServerPort = int64(conn.Addr().(*net.TCPAddr).Port)

	// todo react
}
