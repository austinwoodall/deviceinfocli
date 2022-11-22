package cmd

import (
	"deviceinfocli/utility"
	"fmt"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/spf13/cobra"
)

type HostInfo struct {
	Hostname             string `json:"hostname"`
	Uptime               uint64 `json:"uptime"`
	BootTime             uint64 `json:"bootTime"`
	Procs                uint64 `json:"procs"`
	OS                   string `json:"os"`
	Platform             string `json:"platform"`
	PlatformFamily       string `json:"platformFamily"`
	PlatformVersion      string `json:"platformVersion"`
	KernelVersion        string `json:"kernelVersion"`
	KernelArch           string `json:"kernelArch"`
	VirtualizationSystem string `json:"virtualizationSystem"`
	VirtualizationRole   string `json:"virtualizationRole"`
	HostID               string `json:"hostId"`
}

var hostInfo string

func init() {

	var hostCmd = &cobra.Command{
		Use:     "host",
		Aliases: []string{"h"},
		Long:    `Prints out host info.`,
		Run:     executeHostInfoCmd,
	}

	hostCmd.Flags().StringVarP(&hostInfo, "info", "i", "", "Print specific info about Host instead of all of it.")

	RootCmd.AddCommand(hostCmd)
}

func executeHostInfoCmd(cmd *cobra.Command, args []string) {
	c, _ := host.Info()

	switch hostInfo {
	case "hostname":
		fmt.Printf("HostName: %s", c.Hostname)
		break
	case "uptime":
		fmt.Println(utility.TimeConversion(int(c.Uptime)))
		break
	case "os":
		fmt.Printf("Operating System: %s", c.OS)
		break
	case "kernelv":
		fmt.Printf("Kernel Version: %s", c.KernelVersion)
		break
	case "kernelp":
		fmt.Printf("Kernel Platform: %s", c.KernelArch)
		break
	default:
		fmt.Println(c)
	}

}
