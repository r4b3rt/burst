package client

import (
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/util/jsonutil"
	"github.com/fzdwx/burst/util/log"
	"github.com/lxzan/gws"
	"log/slog"
)

type t struct {
	item *api.PortMappingResp
	cm   *connectionManager
}

func NewTransform(item *api.PortMappingResp) *t {
	return &t{
		item: item,
	}
}

func (t *t) OnOpen(socket *gws.Conn) {
	data := &api.TransFromData{
		ConnectionId: t.item.ConnectionId,
	}

	err := socket.WriteMessage(gws.OpcodeText, jsonutil.Encode(data))
	if err != nil {
		slog.Error("send hello to server error",
			log.ConnectionId(t.item.ConnectionId),
			log.Mapping(t.item.Mapping), log.Reason(err))
		return
	}

	t.cm = &connectionManager{
		connectionId:   t.item.ConnectionId,
		connectionPool: make(map[string]*connection),
	}
}

func (t *t) OnClose(socket *gws.Conn, err error) {
	//TODO implement me
	panic("implement me")
}

func (t *t) OnPing(socket *gws.Conn, payload []byte) {
	//TODO implement me
	panic("implement me")
}

func (t *t) OnPong(socket *gws.Conn, payload []byte) {
	//TODO implement me
	panic("implement me")
}

func (t *t) OnMessage(socket *gws.Conn, message *gws.Message) {
	switch message.Opcode {
	case gws.OpcodeBinary:
		slog.Info("recv data from server",
			log.ClientReadFromServer(),
			log.ConnectionId(t.item.ConnectionId),
			log.Mapping(t.item.Mapping),
		)

		data := jsonutil.Decode(message.Data)
		userConn, err := t.cm.getConnection(data.UserConnectionId, t.item.Mapping)
		if err != nil {
			slog.Error("get connection error",
				log.ConnectionId(t.item.ConnectionId),
				log.Mapping(t.item.Mapping),
				log.UserConnectionId(data.UserConnectionId),
				log.Reason(err))
			return
		}
		if userConn.startRead.Load() == false {
			userConn.startRead.Store(true)
			go localToServer(userConn, socket)
		}

		// 4. send user data to src port
		_, err = userConn.conn.Write(data.Data)
		if err != nil {
			slog.Error("write data to user error",
				log.ClientToLocal(),
				log.ConnectionId(t.item.ConnectionId),
				log.Mapping(t.item.Mapping),
				log.UserConnectionId(data.UserConnectionId),
				log.Reason(err),
			)
			return
		}
		slog.Info("send data to user success",
			log.ClientToLocal(),
			log.ConnectionId(t.item.ConnectionId),
			log.Mapping(t.item.Mapping),
			log.UserConnectionId(data.UserConnectionId))

	}
}

func localToServer(userConn *connection, serverStream *gws.Conn) {
	buf := make([]byte, 1024)
	for {
		// 5. read local data
		n, err := userConn.conn.Read(buf)
		if err != nil {
			slog.Error("read local data, local to server stop",
				log.ClientReadFromLocal(),
				log.ConnectionId(userConn.connectionId),
				log.Mapping(userConn.mapping),
				log.UserConnectionId(userConn.userId),
				log.Reason(err),
			)
			return
		}
		if n == 0 {
			continue
		}
		slog.Info("read local data success",
			log.ClientReadFromLocal(),
			log.ConnectionId(userConn.connectionId),
			log.Mapping(userConn.mapping),
			log.UserConnectionId(userConn.userId))

		// 6. send to server
		if err = serverStream.WriteMessage(gws.OpcodeBinary, jsonutil.Encode(&api.TransFromData{
			ConnectionId:     userConn.connectionId,
			UserConnectionId: userConn.userId,
			Data:             buf[:n],
		})); err != nil {
			slog.Error("send data to server error",
				log.ClientToServer(),
				log.ConnectionId(userConn.connectionId),
				log.Mapping(userConn.mapping),
				log.UserConnectionId(userConn.userId),
				log.Reason(err),
			)
			return
		}
		slog.Info("send data to server success",
			log.ClientToServer(),
			log.ConnectionId(userConn.connectionId),
			log.Mapping(userConn.mapping),
			log.UserConnectionId(userConn.userId))
	}
}
