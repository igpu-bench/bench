/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// resultsCmd represents the results command
var resultsCmd = &cobra.Command{
	Use:   "results",
	Short: "Manage benchmark results",
	Long:  `Commands to manage the results of tests and export results in various formats`,
}

func init() {
	rootCmd.AddCommand(resultsCmd)
}
