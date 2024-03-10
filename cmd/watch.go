package cmd

import (
	"sincro/pkg/watch"

	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch short",
	Long:  `Watch long`,
	Run: func(cmd *cobra.Command, args []string) {
		watch.Start()
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
