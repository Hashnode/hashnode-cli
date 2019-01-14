package cmd

import (
	"github.com/hashnode/hashnode-cli/pkg/posts"
	"github.com/spf13/cobra"
)

//flags
var (
	trending bool
)

// storiesCmd represents the stories command
var storiesCmd = &cobra.Command{
	Use:     "stories",
	Short:   "Read stories published on hashnode",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case trending:
			posts.GetTrendingPosts()
		default:
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(storiesCmd)

	storiesCmd.PersistentFlags().BoolVar(&trending, "hot", false, "get hot trending stories")
}
