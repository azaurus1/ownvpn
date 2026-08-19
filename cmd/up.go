/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vishvananda/netlink"
)

var neighbours map[string]string

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "start ownvpn",
	Run:   Up,
}

func Up(cmd *cobra.Command, args []string) {

	// create TUN
	la := netlink.NewLinkAttrs()
	la.Name = "ownvpn0"
	tun := &netlink.Tuntap{
		LinkAttrs: la,
		Mode:      netlink.TUNTAP_MODE_TUN,
	}

	err := netlink.LinkAdd(tun)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	// assign address
	addr, _ := netlink.ParseAddr("69.255.0.1/32")
	err = netlink.AddrAdd(tun, addr)
	if err != nil {
		log.Println("error: ", err)
		return
	}

	err = netlink.LinkSetUp(tun)
	if err != nil {
		log.Println("error: ", err)
		return
	}

}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
