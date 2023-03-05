/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"github.com/spf13/cobra"
)

var ActiveFiles []string

// samplesCmd represents the samples command
var samplesCmd = &cobra.Command{
	Use:   "samples",
	Short: "Manage sample media files",
	Long:  `Commands to manage sample media files used during transcode testing`,
}

func init() {
	rootCmd.AddCommand(samplesCmd)

}

func generate_active_files_from_sample_dir() error {

	return nil
}
