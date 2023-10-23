package cmd

import (
	"context"
	"fmt"
	"github.com/fzdwx/burst/api"
	"github.com/fzdwx/burst/internal/client"
	"github.com/spf13/cobra"
)

var (
	exportCmd = &cobra.Command{
		Use:     "export [serverAddress:port] -p [port...]",
		Short:   "Export client port to server port",
		Example: `burst export :8000 -p 8080,80`,
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := client.Dial(args[0])
			if err != nil {
				return err
			}

			resp, err := c.Export(context.Background(), &api.ExportRequest{
				ClientPort: ports,
			})

			if err != nil {
				return err
			}

			fmt.Println(resp)

			return nil

		},
	}
	ports   []int64
	address string
)

func init() {
	exportCmd.Flags().Int64SliceVarP(&ports, "port", "p", []int64{}, "port to export")
}
