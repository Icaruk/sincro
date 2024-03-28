package cmd

import (
	"sincro/pkg/watch"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch for changes in any of your sources",
	Long:  `Watch for changes in any of your sources`,
	Run: func(cmd *cobra.Command, args []string) {
		watch.Start()
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
