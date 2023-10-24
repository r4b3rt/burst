package serve

import (
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
		conn.serverSideConn.Accept()
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
