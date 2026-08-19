/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/cobra"
	"github.com/vishvananda/netlink"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the ownvpn service",
	Run:   Run,
}

func Run(cmd *cobra.Command, args []string) {

	if len(neighbours) < 1 {
		fmt.Println("neighbours map is required!")
		return
	}

	var wg sync.WaitGroup

	la := netlink.NewLinkAttrs()
	la.Name = "ownvpn0"
	tun := &netlink.Tuntap{
		LinkAttrs: la,
		Mode:      netlink.TUNTAP_MODE_TUN,
		Flags:     netlink.TUNTAP_ONE_QUEUE | netlink.TUNTAP_NO_PI,
		Queues:    1,
	}

	if err := netlink.LinkAdd(tun); err != nil {
		fmt.Println("error attaching to tun: ", err)
		return
	}

	f := tun.Fds[0]
	defer f.Close()

	wg.Go(func() {
		var buf []byte
		for {
			_, err := tun.Fds[0].Read(buf)
			if err != nil {
				log.Println("error: ", err)
			}

			log.Println("buf: ", buf)
		}

	})

	wg.Wait()
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	runCmd.Flags().StringToStringVar(&neighbours, "neighbours", nil, "key=value pairs")
}
