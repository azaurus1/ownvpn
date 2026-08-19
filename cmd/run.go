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

	if len(neighbours) > 0 {
		fmt.Println("neighbours: ", neighbours)
	}

	var wg sync.WaitGroup

	attrs := netlink.NewLinkAttrs()
	attrs.Name = "tun0"
	tun := &netlink.Tuntap{
		LinkAttrs: attrs,
		Mode:      netlink.TUNTAP_MODE_TUN,
		Queues:    1, // required, or Fds gets closed and thrown away
	}

	if err := netlink.LinkAdd(tun); err != nil {
		log.Fatalf("failed to attach: %v", err)
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

			// log.Println("buf: ", buf)
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
