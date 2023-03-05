/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	format string
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary <run_id>",
	Short: "Output a summary of the benchmark results",
	Long:  "Gather",
	Args:  cobra.ExactArgs(1),
	RunE:  output_summary,
}

func output_summary(cmd *cobra.Command, args []string) error {
	var err error

	// TODO get the summary data

	switch args[0] {
	case "text":
		err = errors.New("TODO export as json")
	// case "json":
	// case "csv":
	default:
		err = errors.New("'" + args[0] + "' format not supported")
	}
	return err
}

func init() {
	resultsCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().StringVarP(&format, "format", "f", "text", "the format to output the summary as [json,text]")
	viper.BindPFlag("output_format", summaryCmd.Flags().Lookup("format"))
}
