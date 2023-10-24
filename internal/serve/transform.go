package serve

import (
	"bufio"
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/util/log"
	"log/slog"
)

func (s *server) Transform(ts api.Burst_TransformServer) error {
	data, err := ts.Recv()
	if err != nil {
		return err
	}

	conn := s.getConnection(data.ConnectionId)
	if conn == nil {
		slog.Error("connection not found", log.ConnectionId(data.ConnectionId))
		return fmt.Errorf("connection %s not found", data.ConnectionId)
	}

	go func() {
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
			go s.transformUserToClient(userConn, ts)
		}
	}()

	for {
		// 7. read data from client stream
		ndata, err := ts.Recv()
		if err != nil {
			slog.Error("read data error, stop transform",
				log.ConnectionId(conn.id),
				log.UserConnectionId(data.UserConnectionId),
				log.Reason(err),
				log.ServerReadFromClient())
			return err
		}

		// 8. send to user conn
		userConn := conn.getUserConn(ndata.UserConnectionId)
		if userConn == nil {
			slog.Error("user conn not found, stop transform",
				log.ConnectionId(conn.id),
				log.UserConnectionId(data.UserConnectionId),
				log.Reason(err),
				log.ServerReadFromClient())
			return fmt.Errorf("user conn %s not found", ndata.UserConnectionId)
		}

		if _, err = userConn.conn.Write(ndata.Data); err != nil {
			slog.Error("write data error, stop transform",
				log.ConnectionId(conn.id),
				log.UserConnectionId(data.UserConnectionId),
				log.Reason(err),
				log.ServerReadFromClient())
			return err
		}

		slog.Info("transform data success",
			log.ConnectionId(conn.id),
			log.UserConnectionId(data.UserConnectionId),
			log.ServerSendToUser())
	}
}

func (s *server) transformUserToClient(userConn *userConnection, clientStream api.Burst_TransformServer) {
	r := bufio.NewReader(userConn.conn)
	buf := make([]byte, 1024)
	for {
		// 1. read user conn data
		n, err := r.Read(buf)
		if err != nil {
			slog.Error("read user data, user to client stop",
				log.ServerReadFromUser(),
				log.ConnectionId(userConn.clientConnectionId),
				log.UserConnectionId(userConn.id),
				log.Reason(err))
			return
		}

		// 2. send data to client stream
		if err = clientStream.Send(&api.TransFromData{
			ConnectionId:     userConn.clientConnectionId,
			UserConnectionId: userConn.id,
			Data:             buf[:n],
		}); err != nil {
			slog.Error("send data error, user to client stop",
				log.ServerToClient(),
				log.ConnectionId(userConn.clientConnectionId),
				log.UserConnectionId(userConn.id),
				log.Reason(err))
			return
		}
	}
}
