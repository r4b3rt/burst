package serve

import (
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/util/jsonutil"
	"github.com/fzdwx/burst/util/log"
	"github.com/lxzan/gws"
	"log/slog"
)

type q struct {
	s *server
}

func (q *q) OnOpen(socket *gws.Conn) {
}

func (q *q) OnClose(socket *gws.Conn, err error) {
	//TODO implement me
	panic("implement me")
}

func (q *q) OnPing(socket *gws.Conn, payload []byte) {
}

func (q *q) OnPong(socket *gws.Conn, payload []byte) {
}

func (q *q) OnMessage(socket *gws.Conn, message *gws.Message) {
	switch message.Opcode {
	case gws.OpcodeText:
		go func() {
			data := jsonutil.Decode(message.Data)
			conn := q.s.getConnection(data.ConnectionId)
			for {
				// accept user connect
				c, err := conn.serverSideConn.Accept()
				if err != nil {
					slog.Error("accept error, stop accept user connect",
						log.ConnectionId(conn.id),
						log.Mapping(conn.mapping),
						log.Reason(err))
					return
				}

				userConn := conn.addUserConn(c)
				slog.Info("accept user connect success", log.ConnectionId(conn.id), log.UserConnectionId(userConn.id))
				go q.s.transformUserToClient(userConn, socket)
			}
		}()
	case gws.OpcodeBinary:
		data := jsonutil.Decode(message.Data)
		conn := q.s.getConnection(data.ConnectionId)
		if conn == nil {
			slog.Error("connection not found", log.ConnectionId(data.ConnectionId))
			return
		}

		userConn := conn.getUserConn(data.UserConnectionId)
		if _, err := userConn.conn.Write(data.Data); err != nil {
			slog.Error("write data error, stop transform",
				log.ServerReadFromClient(),
				log.ConnectionId(conn.id),
				log.UserConnectionId(data.UserConnectionId),
				log.Reason(err),
			)
			return
		}

		slog.Info("transform data success",
			log.ServerSendToUser(),
			log.ConnectionId(conn.id),
			log.UserConnectionId(data.UserConnectionId),
		)
	}
}

func (s *server) transformUserToClient(userConn *userConnection, clientStream *gws.Conn) {
	buf := make([]byte, 1024)
	for {
		// 1. read user conn data
		n, err := userConn.conn.Read(buf)
		if err != nil && err.Error() != "EOF" {
			slog.Error("read user data, user to client stop",
				log.ServerReadFromUser(),
				log.ConnectionId(userConn.clientConnectionId),
				log.UserConnectionId(userConn.id),
				log.Reason(err))
			return
		}
		if n == 0 {
			continue
		}

		slog.Info("read user data success",
			log.ServerReadFromUser(),
			log.ConnectionId(userConn.clientConnectionId),
			log.UserConnectionId(userConn.id))

		// 2. send data to client stream
		if err = clientStream.WriteMessage(gws.OpcodeBinary, jsonutil.Encode(&api.TransFromData{
			ConnectionId:     userConn.clientConnectionId,
			UserConnectionId: userConn.id,
			Data:             buf[:n],
		})); err != nil {
			slog.Error("send data error, user to client stop",
				log.ServerToClient(),
				log.ConnectionId(userConn.clientConnectionId),
				log.UserConnectionId(userConn.id),
				log.Reason(err))
			return
		}

		slog.Info("transform data success",
			log.ServerToClient(),
			log.ConnectionId(userConn.clientConnectionId),
			log.UserConnectionId(userConn.id))
	}
}
