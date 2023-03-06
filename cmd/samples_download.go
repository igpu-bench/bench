/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"errors"
	"fmt"
	"io"

	"github.com/Masterminds/semver/v3"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
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
		return (downloadFactory())(ver)
	},
}

func downloadFactory() func(string) error {
	if viper.GetBool("download.use_http") {
		return httpDownload
	} else {
		return torrentDownload
	}
}

func torrentDownload(ver string) error {
	fmt.Println("in torrentDownloader with version " + ver)

	torrentReader, err := getSampleTorrent(ver)
	if err != nil {
		return err
	}

	fmt.Printf("torrentReader: %v\n", torrentReader)

	return nil
}

func httpDownload(ver string) error {
	return errors.New("downloading via HTTP is not yet supported")
}

func init() {
	samplesCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().String("info_repo", "https://github.com/igpu-bench/samples_info", "The repo used to find sample bundle manifests, torrents, and HTTP sources")
	viper.BindPFlag("download.info_repo", downloadCmd.Flags().Lookup("info_repo"))

	downloadCmd.Flags().Bool("http", false, "switch from the default BitTorrent sample source to an HTTP source")
	viper.BindPFlag("download.use_http", downloadCmd.Flags().Lookup("http"))
}

func getSampleInfoRepo(ver string) (*git.Repository, error) {
	var (
		repo *git.Repository
		err  error
	)

	repo, err = git.Clone(memory.NewStorage(), memfs.New(), &git.CloneOptions{
		URL: viper.GetString("download.info_repo"),
	})
	if err != nil {
		return nil, err
	}

	if ver == "" {
		ver, err = getLatestTag(repo)
		if err != nil {
			return nil, err
		}
		fmt.Println("Using sample bundle version " + ver)
	}

	// Checkout the requested version
	wt, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	newBranch, err := getMatchingTag(repo, ver)
	if err != nil {
		return nil, err
	}
	err = wt.Checkout(&git.CheckoutOptions{Branch: newBranch, Force: true, Keep: false})
	if err != nil {
		return nil, err
	}

	return repo, err
}

func getSampleTorrent(ver string) (io.Reader, error) {
	repo, err := getSampleInfoRepo(ver)
	if err != nil {
		return nil, err
	}

	wt, err := repo.Worktree()
	if err != nil {
		return nil, err
	}

	file, err := wt.Filesystem.Open("ibench_samples.torrent")
	if err != nil {

		// if we couldn't file the torrent file, print all the files that are in the repo root
		fmt.Println("Files in repo root:")
		files, _ := wt.Filesystem.ReadDir(".")
		for _, file := range files {
			fmt.Println("\t" + file.Name())
		}

		return nil, err
	}

	return file, err
}

// returns the tag string which is an exact match or tilde (~) match with
func getMatchingTag(repo *git.Repository, semverTag string) (plumbing.ReferenceName, error) {
	refs, err := repo.Tags()
	if err != nil {
		return plumbing.ReferenceName(""), err
	}

	var refNameSlice []plumbing.ReferenceName
	refs.ForEach(func(r *plumbing.Reference) error {
		refNameSlice = append(refNameSlice, r.Name())
		return nil
	})
	// first check for an exact match
	for _, r := range refNameSlice {
		if util.IsExactVersionMatch(semverTag, r.Short()) {
			return r, nil
		}
	}
	// then check for an fuzzy match
	for _, r := range refNameSlice {
		if util.IsVersionMatch(semverTag, r.Short()) {
			return r, nil
		}
	}
	return plumbing.ReferenceName(""), errors.New("no sample_info repo branch " + semverTag)
}

func getLatestTag(repo *git.Repository) (string, error) {
	refs, err := repo.Tags()
	if err != nil {
		return "", err
	}

	var refSemvers semver.Collection
	refs.ForEach(func(r *plumbing.Reference) error {
		ver, err := semver.NewVersion(r.Name().Short())
		if err == nil {
			refSemvers = append(refSemvers, ver)
		}
		return nil
	})
	// return the largest (most recent) version tag
	return refSemvers[len(refSemvers)-1].Original(), nil
}
