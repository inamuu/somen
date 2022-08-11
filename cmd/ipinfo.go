/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ipinfoCmd = &cobra.Command{
	Use:   "ipinfo",
	Short: "Print the result that reverse ip address information",
	Run: func(cmd *cobra.Command, args []string) {
		ipaddr, _ := cmd.Flags().GetString("ip")
		if ipaddr != "" {
			fmt.Println("IP:", ipaddr)
		} else {
			fmt.Println("Please type ip address as args")
		}
	},
}

func init() {
	rootCmd.AddCommand(ipinfoCmd)
	ipinfoCmd.PersistentFlags().StringP("ip", "i", "", "IP address")
}
