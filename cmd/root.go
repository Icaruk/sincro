package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const __VERSION__ string = "0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sincro",
	Short: "Sincro is a tool to sync your files between a source and multiple destinations",
	Long:  `Sincro is a tool to sync your files between a source and multiple destinations. Where the "source" is the source of truth and the "destinations" are the destinations you want to sync.`,
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sincro.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.Version = string(__VERSION__)
}
