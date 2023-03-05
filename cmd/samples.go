/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// samplesCmd represents the samples command
var samplesCmd = &cobra.Command{
	Use:   "samples",
	Short: "Manage sample media files",
	Long:  `Commands to manage sample media files used during transcode testing`,
}

func init() {
	rootCmd.AddCommand(samplesCmd)
}
