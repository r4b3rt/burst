package serve

import (
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/util/log"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func (s *server) Export(ctx *gin.Context, request *api.ExportRequest) (*api.ExportResponse, error) {
	slog.Info("handle export",
		slog.String("portMapping", toString(request.PortMapping)),
		slog.String("clientAddr", ctx.RemoteIP()),
	)

	var (
		resp = &api.ExportResponse{
			Items: make([]*api.PortMappingResp, 0),
		}
	)

	for i := range request.PortMapping {
		mapping := request.PortMapping[i]
		connectionId, err := s.mapping(mapping)
		if err != nil {
			slog.ErrorContext(ctx, "mapping error", log.Mapping(mapping), log.Reason(err))
		}

		resp.Items = append(resp.Items, &api.PortMappingResp{
			Mapping:      mapping,
			Error:        errToString(err),
			ConnectionId: connectionId,
		})
	}

	return resp, nil
}

func errToString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func toString(mapping []*api.PortMapping) string {
	var str string
	for i, m := range mapping {
		if i != 0 {
			str += ", "
		}
		str += fmt.Sprintf("%d to %d", m.ClientPort, m.ServerPort)
	}
	return str
}
