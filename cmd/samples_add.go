/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <file|directory>",
	Short: "Adds files to the active list of sample files",
	Long: `Adds the given file (or all files in the directory) to the list of
files to use for transcode testing`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init() {
	samplesCmd.AddCommand(addCmd)
}
