/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ipinfoCmd represents the ipinfo command
var ipinfoCmd = &cobra.Command{
	Use:   "ipinfo",
	Short: "Print the result that reverse ip address information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ipinfo called")
	},
}

func init() {
	rootCmd.AddCommand(ipinfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ipinfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	ipinfoCmd.Flags().BoolP("", "i", false, "IP Address")
}
