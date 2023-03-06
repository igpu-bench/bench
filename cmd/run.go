/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	no_download bool
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a benchmark start-to-end",
	Long: `Runs a full benchmark. Begins with downloading samples if there
no active samples. Then runs the benchmark and generates results summaries.`,
	RunE: do_run,
}

func do_run(cmd *cobra.Command, args []string) error {
	var err error
	fmt.Println("TODO do_run")

	if !no_download {
		err = downloadFactory()("")
	}
	return err
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolVarP(&no_download, "no-download", "n", false, "do not download any sample files")
}
