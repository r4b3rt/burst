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
			go s.transformUserToClient(userConn, ts)
		}
	}()

	for {
		var v = &api.TransFromData{}
		err = ts.RecvMsg(v)
		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println(v)
	}
}

func (s *server) transformUserToClient(userConn *userConnection, clientStream api.Burst_TransformServer) {
	r := bufio.NewReader(userConn.conn)
	buf := make([]byte, 1024)
	for {
		// 1. read user conn data
		n, err := r.Read(buf)
		if err != nil {
			slog.Error("read user conn error, stop read",
				log.UserToClient(),
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
			slog.Error("send data to client error, stop send",
				log.UserToClient(),
				log.ConnectionId(userConn.clientConnectionId),
				log.UserConnectionId(userConn.id),
				log.Reason(err))
			return
		}
	}
}
