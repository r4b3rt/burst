package client

import (
	"bufio"
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/util/log"
	"log/slog"
)

func Transform(serverStream api.Burst_TransformClient, m *api.PortMappingResp) {
	data := &api.TransFromData{
		ConnectionId: m.ConnectionId,
	}
	err := serverStream.SendMsg(data)
	if err != nil {
		slog.Error("send hello to server error",
			log.ConnectionId(m.ConnectionId),
			log.Mapping(m.Mapping), log.Reason(err))
		return
	}

	var cm = &connectionManager{
		connectionId:   m.ConnectionId,
		connectionPool: make(map[string]*connection),
	}

	for {
		// 3. recv user data
		data, err = serverStream.Recv()
		if err != nil {
			slog.Error("recv data from server error",
				log.ConnectionId(m.ConnectionId),
				log.Mapping(m.Mapping),
				log.Reason(err),
				log.ClientReadFromServer())
			return
		}
		slog.Info("recv data from server",
			log.ConnectionId(m.ConnectionId),
			log.Mapping(m.Mapping),
			log.ClientReadFromServer())

		userConn, err := cm.getConnection(data.UserConnectionId, m.Mapping)
		if err != nil {
			slog.Error("get connection error",
				log.ConnectionId(m.ConnectionId),
				log.Mapping(m.Mapping),
				log.UserConnectionId(data.UserConnectionId),
				log.Reason(err))
			return
		}
		if userConn.startRead.Load() == false {
			userConn.startRead.Store(true)
			go localToServer(userConn, serverStream)
		}

		// 4. send user data to src port
		_, err = userConn.conn.Write(data.Data)
		if err != nil {
			slog.Error("write data to user error",
				log.ConnectionId(m.ConnectionId),
				log.Mapping(m.Mapping),
				log.UserConnectionId(data.UserConnectionId),
				log.Reason(err),
				log.ClientToLocal())
			return
		}
	}
}

func localToServer(userConn *connection, serverStream api.Burst_TransformClient) {
	reader := bufio.NewReader(userConn.conn)
	buf := make([]byte, 1024)

	for {
		n, err := reader.Read(buf)
		if err != nil {
			slog.Error("read local data, local to server stop",
				log.ConnectionId(userConn.connectionId),
				log.Mapping(userConn.mapping),
				log.UserConnectionId(userConn.userId),
				log.Reason(err),
				log.ClientReadFromLocal())
			return
		}
		if err = serverStream.Send(&api.TransFromData{
			ConnectionId:     userConn.connectionId,
			UserConnectionId: userConn.userId,
			Data:             buf[:n],
		}); err != nil {
			slog.Error("send data to server error",
				log.ConnectionId(userConn.connectionId),
				log.Mapping(userConn.mapping),
				log.UserConnectionId(userConn.userId),
				log.Reason(err),
				log.ClientToServer())
			return
		}
	}
}
