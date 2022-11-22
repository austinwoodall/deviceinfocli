package cmd

import (
	"deviceinfocli/utility"
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/spf13/cobra"
)

var diskInfo string
var path string

func init() {
	var diskCmd = &cobra.Command{
		Use:     "dsk",
		Aliases: []string{"d"},
		Short:   "Print storage usage.",
		Run:     executeStorageCmd,
	}
	diskCmd.Flags().StringVarP(&diskInfo, "info", "i", "", "Print specific info about Disk instead of all of it.")
	diskCmd.Flags().StringVarP(&path, "path", "p", "", "Path to partition.")
	RootCmd.AddCommand(diskCmd)
}

func executeStorageCmd(cmd *cobra.Command, args []string) {

	switch diskInfo {
	case "list":
		disks, _ := disk.Partitions(true)
		for i := 0; i < len(disks); i++ {
			fmt.Println(fmt.Printf("device: %s ", disks[i].Device))
			if i == len(disks) {
				break
			}
		}
		fmt.Println()
		break
	case "usage":
		usage, _ := disk.Usage(path)
		fmt.Println(fmt.Printf("Total: %s", utility.ByteCountIEC(int(usage.Total))))
		fmt.Println(fmt.Printf("Used: %s", utility.ByteCountIEC(int(usage.Used))))
		fmt.Println(fmt.Printf("Free: %s", utility.ByteCountIEC(int(usage.Free))))
		break
	default:
		fmt.Println("disk info")
	}

}
