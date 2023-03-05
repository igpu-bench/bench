/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// benchCmd represents the bench command
var benchCmd = &cobra.Command{
	Use:   "bench",
	Short: "Start a benchmark run",
	Long:  `Begins processing active sample files and records all information into the run's results directory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bench called")
	},
}

func init() {
	rootCmd.AddCommand(benchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// benchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// benchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
