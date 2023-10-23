package cmd

import (
	"github.com/fzdwx/burst/internal/serve"
	"github.com/spf13/cobra"
	"strconv"
)

var (
	serveCmd = &cobra.Command{
		Use:     "serve [port]",
		Short:   "Run the server",
		Example: `burst serve 8000`,
		Args:    cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var (
				err  error
				port int
			)

			if len(args) != 1 {
				port = 8000
			} else {
				port, err = strconv.Atoi(args[0])
				if err != nil {
					return err
				}
			}

			return serve.ListenAndServe(port)
		},
	}
)
