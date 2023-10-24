package serve

import (
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/util/id"
	"net"
	"sync"
)

type connection struct {
	peer net.Addr
	id   string

	serverSideConn net.Listener
	mapping        *api.PortMapping

	userConnections     map[string]*userConnection
	userConnectionsLock sync.RWMutex
}

type userConnection struct {
	id                 string
	clientConnectionId string
	conn               net.Conn
}

func (c *connection) addUserConn(conn net.Conn) *userConnection {
	c.userConnectionsLock.Lock()
	defer c.userConnectionsLock.Unlock()

	var userConn = &userConnection{
		id:                 id.Next(),
		clientConnectionId: c.id,
		conn:               conn,
	}

	c.userConnections[userConn.id] = userConn

	return userConn
}

func (s *server) addConnection(peer net.Addr, m *api.PortMapping, serverSideConn net.Listener) *connection {
	s.connectionLock.Lock()
	defer s.connectionLock.Unlock()

	c := &connection{
		peer:            peer,
		id:              id.Next(),
		serverSideConn:  serverSideConn,
		mapping:         m,
		userConnections: map[string]*userConnection{},
	}
	s.connections[c.id] = c
	return c
}

func (s *server) getConnection(id string) *connection {
	s.connectionLock.RLock()
	defer s.connectionLock.RUnlock()

	return s.connections[id]
}
