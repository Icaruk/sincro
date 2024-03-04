package cmd

import (
	"sincro/pkg/scan"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scans short",
	Long:  `Scans long`,
	Run: func(cmd *cobra.Command, args []string) {
		scan.Start()
	},
}

func init() {
	// rootCmd.AddCommand(pushCmd)
}
