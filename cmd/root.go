package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "infogo",
	Short: "A cli tool to get device information.",
	Run: func(cmd *cobra.Command, args []string) {
		// Empty for now
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
