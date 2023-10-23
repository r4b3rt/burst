package client

import (
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/internal/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type client struct {
	api.BurstClient
}

func Dial(address string, port int) (*client, error) {
	var opts []grpc.DialOption = []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", address, port), opts...)
	if err != nil {
		slog.Error("fail to connect server", log.Reason(err))
		return nil, err
	}

	c := &client{}

	c.BurstClient = api.NewBurstClient(conn)
	return c, nil
}
