package cmd

import (
	"github.com/fzdwx/burst/internal/serve"
	"github.com/spf13/cobra"
)

var (
	serveCmd = &cobra.Command{
		Use:   "serve",
		Short: "Run the server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return serve.ListenAndServe(cmd, args)
		},
	}
)
