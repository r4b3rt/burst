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
		Use:   "export",
		Short: "Export client port to server port",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := client.Dial("localhost", 8000)
			if err != nil {
				return err
			}

			resp, err := c.Export(context.Background(), &api.ExportRequest{
				ClientPort: 8888,
			})

			if err != nil {
				return err
			}

			fmt.Println(resp)

			return nil

		},
	}
)
