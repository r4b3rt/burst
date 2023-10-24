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

func UserConnectionId(id string) slog.Attr {
	return slog.String("userConnectionId", id)
}

func Mapping(m *api.PortMapping) slog.Attr {
	portMappingStr := fmt.Sprintf("%d to %d", m.ClientPort, m.ServerPort)
	return slog.String("port mapping", portMappingStr)
}

func ServerReadFromUser() slog.Attr {
	return slog.String("direction", "server <- user")
}

func ServerToClient() slog.Attr {
	return slog.String("direction", "server -> client")
}

func ClientReadFromServer() slog.Attr {
	return slog.String("direction", "client <- server")
}

func ClientToLocal() slog.Attr {
	return slog.String("direction", "client -> local")
}

func ClientReadFromLocal() slog.Attr {
	return slog.String("direction", "client <- local")
}

func ClientToServer() slog.Attr {
	return slog.String("direction", "client -> server")
}
