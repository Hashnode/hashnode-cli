package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

const (
	rootAPIURL = "https://hashnode.com/ajax/posts"
)

// API URL's
var (
	hotPostsAPI = fmt.Sprintf("%s/%s", rootAPIURL, "hot")
	newsAPI     = fmt.Sprintf("%s/%s", rootAPIURL, "news")
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
			getNews()
		case hot:
			getHotPosts()
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

func getHotPosts() {
	b, err := makeRequest(hotPostsAPI)
	if err != nil {
		log.Println(err)
	}
	var hotposts HotPosts
	err = json.Unmarshal(b, &hotposts)
	if err != nil {
		log.Println(err)
	}
	var posttitles []string
	for _, post := range hotposts.Posts {
		posttitles = append(posttitles, post.Title)
	}
	prompt := promptui.Select{
		Label: "Hot Posts",
		Items: posttitles,
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
	box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}

}

func getNews() {
	b, err := makeRequest(newsAPI)
	if err != nil {
		log.Println(err)
	}
	m := make(map[string]interface{})
	json.Unmarshal(b, &m)
	fmt.Println(m)
}

func makeRequest(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Duration(1 * time.Minute),
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		// read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return body, nil
	}

	return nil, fmt.Errorf("Not Found")

}

type HotPosts struct {
	Posts []struct {
		ID             string `json:"_id"`
		FollowersCount int    `json:"followersCount"`
		Author         struct {
			ID            string      `json:"_id"`
			Role          interface{} `json:"role"`
			NumFollowing  int         `json:"numFollowing"`
			NumFollowers  int         `json:"numFollowers"`
			Name          string      `json:"name"`
			Tagline       string      `json:"tagline"`
			Photo         string      `json:"photo"`
			Username      string      `json:"username"`
			Appreciations []struct {
				Badge string `json:"badge"`
				ID    string `json:"_id"`
				Count int    `json:"count"`
			} `json:"appreciations"`
			DateJoined  time.Time `json:"dateJoined"`
			SocialMedia struct {
				Website  string `json:"website"`
				Twitter  string `json:"twitter"`
				Github   string `json:"github"`
				Linkedin string `json:"linkedin"`
				Google   string `json:"google"`
				Facebook string `json:"facebook"`
			} `json:"socialMedia"`
			StoriesCreated       []string      `json:"storiesCreated"`
			Location             string        `json:"location"`
			CoverImage           string        `json:"coverImage"`
			BadgesAwarded        []interface{} `json:"badgesAwarded"`
			TotalUpvotesReceived int           `json:"totalUpvotesReceived"`
			IsEvangelist         bool          `json:"isEvangelist"`
			NumReactions         int           `json:"numReactions"`
		} `json:"author"`
		Cuid                   string        `json:"cuid"`
		Slug                   string        `json:"slug"`
		Title                  string        `json:"title"`
		Type                   string        `json:"type"`
		ReactionsByCurrentUser []interface{} `json:"reactionsByCurrentUser"`
		TotalReactions         int           `json:"totalReactions"`
		Reactions              []struct {
			ID    string `json:"_id"`
			Image string `json:"image"`
			Name  string `json:"name"`
		} `json:"reactions"`
		BookmarkedIn []interface{} `json:"bookmarkedIn"`
		HasReward    bool          `json:"hasReward"`
		Contributors []struct {
			User struct {
				ID            string        `json:"_id"`
				Username      string        `json:"username"`
				Name          string        `json:"name"`
				Photo         string        `json:"photo"`
				Tagline       string        `json:"tagline"`
				BadgesAwarded []interface{} `json:"badgesAwarded"`
				Appreciations []struct {
					Badge string `json:"badge"`
					ID    string `json:"_id"`
					Count int    `json:"count"`
				} `json:"appreciations"`
				DateJoined  time.Time `json:"dateJoined"`
				SocialMedia struct {
					Twitter       string `json:"twitter"`
					Github        string `json:"github"`
					Stackoverflow string `json:"stackoverflow"`
					Linkedin      string `json:"linkedin"`
					Google        string `json:"google"`
					Website       string `json:"website"`
				} `json:"socialMedia"`
				StoriesCreated       []string    `json:"storiesCreated"`
				NumFollowing         int         `json:"numFollowing"`
				NumFollowers         int         `json:"numFollowers"`
				Location             string      `json:"location"`
				Role                 interface{} `json:"role"`
				CoverImage           string      `json:"coverImage"`
				TotalUpvotesReceived int         `json:"totalUpvotesReceived"`
				IsEvangelist         bool        `json:"isEvangelist"`
				NumReactions         int         `json:"numReactions"`
			} `json:"user"`
			Stamp string `json:"stamp"`
			ID    string `json:"_id"`
		} `json:"contributors"`
		IsActive      bool      `json:"isActive"`
		ResponseCount int       `json:"responseCount"`
		DateAdded     time.Time `json:"dateAdded"`
		Tags          []struct {
			ID         string      `json:"_id"`
			Name       string      `json:"name"`
			Slug       string      `json:"slug"`
			MergedWith interface{} `json:"mergedWith,omitempty"`
			IsApproved bool        `json:"isApproved"`
			IsActive   bool        `json:"isActive"`
		} `json:"tags"`
		Downvotes               int           `json:"downvotes"`
		Upvotes                 int           `json:"upvotes"`
		TotalPollVotes          int           `json:"totalPollVotes"`
		PollOptions             []interface{} `json:"pollOptions"`
		HasPolls                bool          `json:"hasPolls"`
		Brief                   string        `json:"brief"`
		CoverImage              string        `json:"coverImage"`
		Views                   int           `json:"views"`
		IsAnonymous             bool          `json:"isAnonymous"`
		DateUpdated             time.Time     `json:"dateUpdated"`
		IndexVotedByCurrentUser int           `json:"indexVotedByCurrentUser"`
		IsFollowing             bool          `json:"isFollowing"`
		DateFeatured            time.Time     `json:"dateFeatured,omitempty"`
	} `json:"posts"`
}
