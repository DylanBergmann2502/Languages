/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	client  = http.Client{
		Timeout: 2 * time.Second,
	}
)

func ping(domain string) (int, error) {
	url := "http://" + domain
	request, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}
	response.Body.Close()
	return response.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if resp, err := ping(urlPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "The URL to ping")
	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}
	NetCmd.AddCommand(pingCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
