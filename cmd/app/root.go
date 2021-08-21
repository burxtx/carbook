package app

import (
	"flag"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "user",
		Short: "user server",
	}
)

func Execute() error {
	rootCmd.Flags().AddGoFlagSet(flag.CommandLine)
	return rootCmd.Execute()
}
