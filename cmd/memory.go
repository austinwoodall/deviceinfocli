package cmd

import (
	"deviceinfocli/utility"
	"fmt"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
)

var memInfo string

func init() {
	var ramCmd = &cobra.Command{
		Use:     "mem",
		Aliases: []string{"m"},
		Short:   "Ram usage.",
		Long:    `Prints out the device ram usage.`,
		Run:     executeMemoryCmd,
	}
	ramCmd.Flags().StringVarP(&memInfo, "info", "i", "", "Print specific info about Memory instead of all of it.")
	RootCmd.AddCommand(ramCmd)
}

func executeMemoryCmd(cmd *cobra.Command, args []string) {
	v, _ := mem.VirtualMemory()

	switch memInfo {
	case "total":
		fmt.Printf("%s Total", utility.ByteCountIEC(int(v.Total)))
		break
	case "free":
		fmt.Printf("%s Total", utility.ByteCountIEC(int(v.Free)))
		break
	case "used":
		fmt.Printf("%s Total", utility.ByteCountIEC(int(v.Used)))
		break
	case "cached":
		fmt.Printf("%s Total", utility.ByteCountIEC(int(v.Cached)))
		break
	default:
		fmt.Println(v)
	}

}
