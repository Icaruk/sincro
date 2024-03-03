package cmd

import (
	"sincro/pkg/push"

	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "pushs short",
	Long:  `pushs long`,
	Run: func(cmd *cobra.Command, args []string) {
		push.Local()
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
}
