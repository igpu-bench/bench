/*
Copyright © 2023 iGPU Bench Team
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ibench",
	Short: "ibench is a iGPU transcoding benchmark tool",
	Long: `ibench benchmarks integrated GPU performance in media transcode.
ibench has its own sample media or can use provided samples. Results
are collected with summaries making it easy to compare performance
across different iGPUs.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if debug, err := cmd.Flags().GetBool("debug"); debug && err == nil {
			fmt.Println("Flags:")
			cmd.Flags().VisitAll(func(f *pflag.Flag) { fmt.Println("\t" + f.Name + ": " + f.Value.String()) })
			fmt.Println("Config:")
			for k, v := range viper.AllSettings() {
				fmt.Printf("\t%s: %v\n", k, v)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ibench.yaml)")
	rootCmd.PersistentFlags().String("samples_dir", "samples/", "Path to the location for downloading samples")
	rootCmd.PersistentFlags().String("transcode_dir", "transcode/", "Path to the location for outputting transcoded files")
	rootCmd.PersistentFlags().String("results_dir", "results/", "Path to the location for downloading samples")
	rootBinder("samples_dir")
	rootBinder("transcode_dir")
	rootBinder("results_dir")

	rootCmd.PersistentFlags().Bool("debug", false, "enables extra logging and debug information")
	rootCmd.PersistentFlags().CountP("verbose", "v", "enables more verbose ")
	rootBinder("debug")
	rootBinder("verbose")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.SetConfigName("ibench")
		viper.SetConfigType("yaml")

		viper.AddConfigPath("/config/")
		viper.AddConfigPath(home)
	}

	viper.SetEnvPrefix("ib")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

func rootBinder(name string) error {
	return viper.BindPFlag(name, rootCmd.PersistentFlags().Lookup(name))
}
