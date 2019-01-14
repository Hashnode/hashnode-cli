package cmd

import (
	"github.com/hashnode/hashnode-cli/pkg/posts"
	"github.com/spf13/cobra"
)

// flags
var (
	hot bool
)

// postsCmd represents the posts command
var postsCmd = &cobra.Command{
	Use:     "dicussions",
	Short:   "Read discussions on hashnode",
	Aliases: []string{"d"},
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case hot:
			posts.GetHotPosts()
		default:
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(postsCmd)

	postsCmd.PersistentFlags().BoolVar(&hot, "hot", false, "get hot posts")
}
