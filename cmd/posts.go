package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// postsCmd represents the posts command
var postsCmd = &cobra.Command{
	Use:   "posts",
	Short: "Lists posts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("posts called")
	},
}

func init() {
	rootCmd.AddCommand(postsCmd)
}
