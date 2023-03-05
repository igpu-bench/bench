/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove <file|directory>",
	Short: "Remove files from the active list of sample files",
	Long: `Removes the given file (or all files in the directory) from the list of
files to use for transcode testing`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("remove called")
	},
}

func init() {
	samplesCmd.AddCommand(removeCmd)

	removeCmd.Flags().Bool("delete", false, "delete the sample file in addition to removing from active files")
}
