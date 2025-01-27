/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Sends a DELETE request to the specified URL",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		client := &http.Client{}

		req, err := http.NewRequest("DELETE", url, nil) // Changed to DELETE method
		if err != nil {
			fmt.Println("Error creating request: ", err)
			return
		}

		// Set the headers
		for _, header := range headers {
			k, v := parseHeader(header)
			req.Header.Set(k, v)
		}

		// Set Auth header
		if auth != "" {
			if strings.Contains(auth, ":") {
				parts := strings.SplitN(auth, ":", 2)
				req.SetBasicAuth(parts[0], parts[1])
			} else {
				req.Header.Set("Authorization", "Bearer "+auth) // Added space to the bearer token
			}
		}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request: ", err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response: ", err)
			return
		}

		fmt.Printf("Status: %s\n", resp.Status)
		fmt.Println("Body: ")
		fmt.Println(string(pretty.Pretty(body)))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&url, "url", "u", "", "Target URL (required)")
	deleteCmd.Flags().StringSliceVarP(&headers, "header", "H", []string{}, "Headers (key:value)")
	deleteCmd.Flags().StringVarP(&auth, "auth", "a", "", "Auth (user:pass or token)")
	deleteCmd.MarkFlagRequired("url")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
