package posts

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/rivo/tview"
)

var (
	trendingStoriesAPI = fmt.Sprintf("%s/%s/%s/%s", rootAPIURL, "posts", "stories", "trending")
)

func GetTrendingPosts() {
	b, err := makeRequest(trendingStoriesAPI)
	if err != nil {
		log.Printf("Oops, some network error: %v\n", err)
		os.Exit(0)
	}
	var posts TrendingStories
	err = json.Unmarshal(b, &posts)
	if err != nil {
		log.Println(err)
	}

	list := tview.NewList()
	app := tview.NewApplication()

	for ind, post := range posts.Posts {

		list = list.AddItem(post.Title, post.Brief, rune(strconv.Itoa(ind)[0]), nil)
	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
		os.Exit(0)
	})

	list.SetSelectedFunc(func(runeindex int, title string, desc string, r rune) {
		if r != 'q' {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				app.Stop()
				panic(err)
			}

			openPost(app, posts.Posts[n].Cuid, list)

		}
	})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		app.Stop()
		panic(err)
	}

}

// Types

type TrendingStories struct {
	Posts []struct {
		ID             string `json:"_id"`
		FollowersCount int    `json:"followersCount"`
		Author         struct {
			ID       string `json:"_id"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"author"`
		Cuid           string        `json:"cuid"`
		Slug           string        `json:"slug"`
		Title          string        `json:"title"`
		Type           string        `json:"type"`
		TotalReactions int           `json:"totalReactions"`
		BookmarkedIn   []interface{} `json:"bookmarkedIn"`
		HasReward      bool          `json:"hasReward"`
		IsPublication  bool          `json:"isPublication,omitempty"`
		Contributors   []struct {
			User  string `json:"user"`
			Stamp string `json:"stamp"`
			ID    string `json:"_id"`
		} `json:"contributors"`
		IsActive       bool          `json:"isActive"`
		ResponseCount  int           `json:"responseCount"`
		DateAdded      time.Time     `json:"dateAdded"`
		Tags           []string      `json:"tags"`
		Downvotes      int           `json:"downvotes"`
		Upvotes        int           `json:"upvotes"`
		TotalPollVotes int           `json:"totalPollVotes"`
		PollOptions    []interface{} `json:"pollOptions"`
		HasPolls       bool          `json:"hasPolls"`
		Brief          string        `json:"brief"`
		CoverImage     string        `json:"coverImage"`
		Views          int           `json:"views"`
		IsAnonymous    bool          `json:"isAnonymous"`
		DateUpdated    time.Time     `json:"dateUpdated,omitempty"`
		IsOriginal     bool          `json:"isOriginal,omitempty"`
		DateFeatured   time.Time     `json:"dateFeatured,omitempty"`
	} `json:"posts"`
}
