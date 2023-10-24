package client

import (
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/internal/cons"
	"github.com/fzdwx/burst/util/log"
	"log/slog"
	"net"
	"sync"
	"sync/atomic"
)

type connectionManager struct {
	connectionId string

	connectionPool     map[string]*connection
	connectionPoolLock sync.RWMutex
}

type connection struct {
	connectionId string
	userId       string
	mapping      *api.PortMapping

	conn      net.Conn
	protocol  string
	startRead atomic.Bool
}

func (cm *connectionManager) getConnection(userId string, m *api.PortMapping) (*connection, error) {
	// 1. get connection from pool
	cm.connectionPoolLock.RLock()
	conn, ok := cm.connectionPool[userId]
	cm.connectionPoolLock.RUnlock()
	if ok {
		return conn, nil
	}

	var newConn = &connection{
		connectionId: cm.connectionId,
		mapping:      m,
		userId:       userId,
		startRead:    atomic.Bool{},
	}

	// 2. if not exist, create connection
	switch m.Protocol {
	case cons.Protocol.TCP:
		dial, err := net.Dial(cons.Protocol.TCP, fmt.Sprintf("%s:%d", "", m.ClientPort))
		if err != nil {
			slog.Error("dial tcp error",
				log.Mapping(m),
				log.ConnectionId(cm.connectionId),
				log.UserConnectionId(userId),
				log.Reason(err))
			return nil, err
		}
		newConn.conn = dial
		newConn.protocol = cons.Protocol.TCP
	default:
		return nil, fmt.Errorf("not support protocol %s", m.Protocol)
	}

	cm.connectionPoolLock.Lock()
	cm.connectionPool[userId] = newConn
	cm.connectionPoolLock.Unlock()

	return newConn, nil
}

func (cm *connectionManager) removeConnection(id string) {
	cm.connectionPoolLock.Lock()
	delete(cm.connectionPool, id)
	cm.connectionPoolLock.Unlock()
}
