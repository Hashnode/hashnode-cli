package posts

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/rivo/tview"
)

var (
	newsAPI = fmt.Sprintf("%s/posts/%s", rootAPIURL, "news")
)

func GetNews() {

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	s.Start()
	b, err := makeRequest(newsAPI)
	if err != nil {
		log.Printf("Oops, some network error: %v\n", err)
		os.Exit(0)
	}
	s.Stop()

	var posts TopNews
	err = json.Unmarshal(b, &posts)
	if err != nil {
		log.Println(err)
	}

	list := tview.NewList()
	list.Box.SetBorder(true)

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
			list.Box.Blur()
			openPost(app, posts.Posts[n].Cuid, list)

		}
	})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		app.Stop()
		panic(err)
	}

}

type TopNews struct {
	Posts []struct {
		ID             string `json:"_id"`
		IsFollowing    bool   `json:"isFollowing"`
		FollowersCount int    `json:"followersCount"`
		Cuid           string `json:"cuid"`
		Slug           string `json:"slug"`
		Author         struct {
			ID                   string        `json:"_id"`
			Username             string        `json:"username"`
			Name                 string        `json:"name"`
			Photo                string        `json:"photo"`
			Tagline              string        `json:"tagline"`
			IsEvangelist         bool          `json:"isEvangelist"`
			BadgesAwarded        []interface{} `json:"badgesAwarded"`
			TotalUpvotesReceived int           `json:"totalUpvotesReceived"`
			Appreciations        []struct {
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
			StoriesCreated []interface{} `json:"storiesCreated"`
			NumFollowing   int           `json:"numFollowing"`
			NumFollowers   int           `json:"numFollowers"`
			IsDeactivated  bool          `json:"isDeactivated"`
			Location       string        `json:"location"`
			NumReactions   int           `json:"numReactions"`
		} `json:"author"`
		Title                  string        `json:"title"`
		URL                    string        `json:"url"`
		Type                   string        `json:"type"`
		Host                   string        `json:"host"`
		ReactionsByCurrentUser []interface{} `json:"reactionsByCurrentUser"`
		TotalReactions         int           `json:"totalReactions"`
		Reactions              []struct {
			ID    string `json:"_id"`
			Image string `json:"image"`
			Name  string `json:"name"`
		} `json:"reactions"`
		BookmarkedIn  []interface{} `json:"bookmarkedIn"`
		HasReward     bool          `json:"hasReward"`
		IsPublication bool          `json:"isPublication"`
		Contributors  []interface{} `json:"contributors"`
		IsActive      bool          `json:"isActive"`
		ResponseCount int           `json:"responseCount"`
		DateAdded     time.Time     `json:"dateAdded"`
		Tags          []struct {
			ID         string      `json:"_id"`
			Name       string      `json:"name"`
			Slug       string      `json:"slug"`
			IsApproved bool        `json:"isApproved"`
			IsActive   bool        `json:"isActive"`
			MergedWith interface{} `json:"mergedWith,omitempty"`
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
		IndexVotedByCurrentUser int           `json:"indexVotedByCurrentUser"`
		OriginalURL             string        `json:"originalUrl"`
	} `json:"posts"`
}
