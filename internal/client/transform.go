package client

import (
	"fmt"
	"github.com/fzdwx/burst/api"
)

func Transform(c api.Burst_TransformClient, mapping []*api.PortMappingResp) {
	m := mapping[0]
	data := &api.TransFromData{
		ConnectionId: m.ConnectionId,
	}
	err := c.SendMsg(data)
	fmt.Println(err)

	c.Recv()
}
