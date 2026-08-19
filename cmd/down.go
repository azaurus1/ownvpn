/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vishvananda/netlink"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "destroy ownvpn interface",
	Run:   Down,
}

func Down(cmd *cobra.Command, args []string) {
	// create TUN
	la := netlink.NewLinkAttrs()
	la.Name = "ownvpn0"
	tun := &netlink.Tuntap{
		LinkAttrs: la,
		Mode:      netlink.TUNTAP_MODE_TUN,
	}
	err := netlink.LinkDel(tun)
	if err != nil {
		log.Println("error: ", err)
		return
	}
}

func init() {
	rootCmd.AddCommand(downCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
