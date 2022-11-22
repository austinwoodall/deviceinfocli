package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/spf13/cobra"
)

func init() {
	info, _ := cpu.Info()
	var cpuCmd = &cobra.Command{
		Use:     "cpu",
		Aliases: []string{"c"},
		Short:   "CPU usage.",
		Long:    `Prints out the device CPU.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
			fmt.Println(info)
		},
	}

	RootCmd.AddCommand(cpuCmd)
}
