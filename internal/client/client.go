package client

import (
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/util/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type client struct {
	api.BurstClient
}

func Dial(address string) (*client, error) {
	var opts = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		slog.Error("fail to connect server", log.Reason(err))
		return nil, err
	}

	c := &client{}

	c.BurstClient = api.NewBurstClient(conn)
	return c, nil
}
