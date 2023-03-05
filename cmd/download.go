/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/igpu-bench/ibench/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download [version]",
	Short: "downloads the latest (or specific) sample bundle",
	Long: "downloads a sample bundle for the latest (or a specific version if passed).\n" +
		"If `--http` is passed, will attempt to use HTTP to download instead of the default bittorrent." +
		"version is semver string (can be prepended by `v`).",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error = nil

		var ver string

		if len(args) == 1 {
			ver = args[0]
			if !util.IsValidVersion(ver) {
				return errors.New("version must be in valid semver format")
			}
		}

		if viper.GetBool("http_download") {
			err = httpDownload(ver)
		} else {
			err = torrentDownload(ver)
		}
		return err
	},
}

func torrentDownload(ver string) error {
	fmt.Println("in torrentDownloader with version " + ver)
	return nil
}

func httpDownload(ver string) error {
	return errors.New("downloading via HTTP is not yet supported")
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	downloadCmd.Flags().Bool("http", false, "switch from the default BitTorrent sample source to an HTTP source")
	viper.BindPFlag("http_download", downloadCmd.Flags().Lookup("http"))
}
