/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

var (
	url     string
	headers []string
	auth    string
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Sends a GET request to the specified URL",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command.,
	Run: func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		for _, header := range headers {
			key, value := parseHeader(header)
			req.Header.Set(key, value)
		}

		if auth != "" {
			if strings.Contains(auth, ":") {
				parts := strings.SplitN(auth, ":", 2)
				req.SetBasicAuth(parts[0], parts[1])
			} else {
				req.Header.Set("Authorization", "Bearer "+auth)
			}
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body: ", err)
			return
		}

		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Printf("Headers: ")
		for k, v := range resp.Header {
			fmt.Printf(" %s: %s\n", k, strings.Join(v, ", "))
		}
		fmt.Println("\nBody:")
		fmt.Println(string(pretty.Pretty(body)))
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&url, "url", "u", "", "Target URL (required)")
	getCmd.Flags().StringSliceVarP(&headers, "header", "H", []string{}, "Headers (key:value)")
	getCmd.Flags().StringVarP(&auth, "auth", "a", "", "Authentication (username:password or Bearer token)")
	getCmd.MarkFlagRequired("url")
}

func parseHeader(header string) (string, string) {
	parts := strings.SplitN(header, ":", 2)
	if len(parts) != 2 {
		return "", ""
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}
