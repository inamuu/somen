/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print th version number of somen",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	//ToDo
	//rootCmd.Flags().BoolP("version", "v", false, "Print version")
}
