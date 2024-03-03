package cmd

import (
	"fmt"
	"sincro/pkg/utils/config"

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
			fmt.Println(reason)
		} else {
			fmt.Println(reason)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
