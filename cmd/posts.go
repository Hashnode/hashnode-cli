package cmd

import (
	"log"

	"github.com/hashnode/hashnode-cli/pkg/posts"
	"github.com/spf13/cobra"
)

// flags
var (
	hot  bool
	news bool
)

// postsCmd represents the posts command
var postsCmd = &cobra.Command{
	Use:   "posts",
	Short: "Lists posts",
	Run: func(cmd *cobra.Command, args []string) {
		switch {
		case news:
			posts.GetNews()
		case hot:
			posts.GetHotPosts()
		default:
			log.Println("Specify what posts to get")
		}
	},
}

func init() {
	rootCmd.AddCommand(postsCmd)

	postsCmd.PersistentFlags().BoolVar(&hot, "hot", false, "get hot posts")
	postsCmd.PersistentFlags().BoolVar(&news, "news", false, "get news")
}
