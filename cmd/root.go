package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	Version = "0.0.1-dev"

	rootCmd = &cobra.Command{
		Use:     "burst",
		Short:   "Expose ports to the server quickly.",
		Version: Version,
	}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
