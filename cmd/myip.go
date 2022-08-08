package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// myipCmd represents the myip command
var myipCmd = &cobra.Command{
	Use:   "myip",
	Short: "Print my global ip address",
	Run: func(cmd *cobra.Command, args []string) {

		url := "https://ifconfig.me/"
		req, _ := http.NewRequest(http.MethodGet, url, nil)

		client := new(http.Client)
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		defer resp.Body.Close()

		//debug
		//fmt.Printf("%-v", resp)

		body, _ := io.ReadAll(resp.Body)

		fmt.Println("MyGlobalIP:", string(body))

	},
}

func init() {
	rootCmd.AddCommand(myipCmd)
}
