package cmd

import (
	"sincro/pkg/utils/config"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates the sincro config file",
	Long:  `Creates the sincro config file (sincro.json) in the root directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		reason, success := config.Init()

		if success {
			color.Green(reason)
		} else {
			color.Red(reason)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
