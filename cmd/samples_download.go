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
	Short: "Downloads the latest (or specific) sample bundle",
	Long: "Downloads a sample bundle for the latest (or a specific version if passed).\n" +
		"If `--http` is passed, will attempt to use HTTP to download instead of the default bittorrent." +
		"version is semver string (can be prepended by `v`).",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var ver string

		if len(args) == 1 {
			ver = args[0]
			if !util.IsValidVersion(ver) {
				return errors.New("version must be in valid semver format")
			}
		}

		// call the correct downloader based on which method is requested
		return (download_factory())(ver)
	},
}

func download_factory() func(string) error {
	if viper.GetBool("download.use_http") {
		return httpDownload
	} else {
		return torrentDownload
	}
}

func torrentDownload(ver string) error {
	fmt.Println("in torrentDownloader with version " + ver)

	return nil
}

func httpDownload(ver string) error {
	return errors.New("downloading via HTTP is not yet supported")
}

func init() {
	samplesCmd.AddCommand(downloadCmd)

	downloadCmd.PersistentFlags().StringP("info_repo", "i", "https://github.com/igpu-bench/samples_info", "The repo used to find sample bundle manifests, torrents, and HTTP sources")
	viper.BindPFlag("download.info_repo", downloadCmd.Flags().Lookup("info_repo"))

	downloadCmd.Flags().Bool("http", false, "switch from the default BitTorrent sample source to an HTTP source")
	viper.BindPFlag("download.use_http", downloadCmd.Flags().Lookup("http"))
}
