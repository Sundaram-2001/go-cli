/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "url-monitor",
	Short: "url-monitor",
	Long:  `url-monitor`,
}

func init() {
	rootCmd.AddCommand(checkUrlCmd, checkStatusCmd)
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
