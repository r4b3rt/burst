package log

import (
	"fmt"
	"github.com/fzdwx/burst/api"
	"log/slog"
)

func Reason(err error) slog.Attr {
	return slog.String("reason", err.Error())
}

func ConnectionId(id string) slog.Attr {
	return slog.String("connectionId", id)
}

func Mapping(m *api.PortMapping) slog.Attr {
	portMappingStr := fmt.Sprintf("%d to %d", m.ClientPort, m.ServerPort)
	return slog.String("port mapping", portMappingStr)
}
