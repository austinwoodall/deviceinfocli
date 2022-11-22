package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Version string

func init() {
	var versionCmd = &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		Short:   "Print application version.",
		Long:    `Software version 1.0`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Version 1.0")
		},
	}
	RootCmd.AddCommand(versionCmd)
}
