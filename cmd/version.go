package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version, commithash string

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of hashnode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Release: %s\nCommit: %s\n", version, commithash)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
