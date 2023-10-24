package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/internal/client"
	"github.com/fzdwx/burst/util/log"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	log2 "log"
	"log/slog"
	"strconv"
	"strings"
)

func init() {
	exportCmd.Flags().StringArrayVarP(&portMapping, "port", "p", []string{}, "port mapping, eg: 8888:18888")
}

var (
	exportCmd = &cobra.Command{
		Use:     "export serverAddress:port -p clientPort:serverPort ...",
		Short:   "Export client port to server port",
		Example: `burst export :8000 -p 8888:18888 -p 9999:19999`,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.Dial(args[0])
			if err != nil {
				log2.Fatal(err)
			}

			mapping, err := validatePortMapping(portMapping)
			if err != nil {
				log2.Fatal(err)
			}

			resp, err := c.Export(context.Background(), &api.ExportRequest{
				PortMapping: mapping,
			})
			if err != nil {
				log2.Fatal(err)
			}

			items := resp.GetItems()
			if items == nil || len(items) == 0 {
				log2.Fatal(fmt.Errorf("not port mapping success"))
			}

			successMapping := lo.Filter(items, func(item *api.PortMappingResp, index int) bool {
				ok := item.Error == ""
				if ok == false {
					slog.Error("mapping error", log.Mapping(item.Mapping), log.Reason(errors.New(item.Error)))
				}
				return ok
			})

			if successMapping == nil || len(successMapping) == 0 {
				log2.Fatal(fmt.Errorf("not port mapping success"))
			}

			transformClient, err := c.Transform(context.Background())
			if err != nil {
				log2.Fatal(err)
			}

			client.Transform(transformClient, successMapping)
		},
	}
	portMapping []string
)

func validatePortMapping(portMapping []string) ([]*api.PortMapping, error) {
	if len(portMapping) == 0 {
		return nil, fmt.Errorf("port mapping is empty")
	}

	var mappings []*api.PortMapping
	for _, p := range portMapping {
		ports := strings.Split(p, ":")
		if len(ports) == 0 {
			return nil, fmt.Errorf("invalid port mapping: %s", p)
		}

		mapping := parse(ports)
		if mapping == nil {
			return nil, fmt.Errorf("invalid port mapping: %s", p)
		}
		mappings = append(mappings, mapping)
	}
	return mappings, nil
}

func parse(ports []string) *api.PortMapping {
	var m = &api.PortMapping{
		Protocol: "tcp", // todo
	}
	m.ClientPort, _ = strconv.ParseInt(ports[0], 10, 64)

	if len(ports) == 1 {
		slog.Info("server port is not set, will random generate", slog.Int64("clientPort", m.ClientPort))
	} else {
		m.ServerPort, _ = strconv.ParseInt(ports[1], 10, 64)
	}

	if m.ClientPort == 0 {
		slog.Error("invalid client port", slog.String("clientPort", ports[0]))
		return nil
	}

	return m
}
