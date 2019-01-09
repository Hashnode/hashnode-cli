package posts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/rivo/tview"
)

const (
	rootAPIURL = "https://hashnode.com/ajax"
)

// API URL's
var (
	hotPostsAPI = fmt.Sprintf("%s/posts/%s", rootAPIURL, "hot")
	postAPI     = fmt.Sprintf("%s/post", rootAPIURL)
)

func GetHotPosts() {
	b, err := makeRequest(hotPostsAPI)
	if err != nil {
		log.Printf("Oops, some network error: %v\n", err)
		os.Exit(0)
	}
	var hotposts HotPosts
	err = json.Unmarshal(b, &hotposts)
	if err != nil {
		log.Println(err)
	}

	list := tview.NewList()
	app := tview.NewApplication()

	for ind, post := range hotposts.Posts {

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

			openPost(app, hotposts.Posts[n].Cuid, list)

		}
	})
	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		app.Stop()
		panic(err)
	}

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

// Types

type HotPosts struct {
	Posts []PostDetails `json:"posts,omitempty"`
}

type PostDetails struct {
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
}

type Post struct {
	Post struct {
		IndexVotedByCurrentUser int        `json:"indexVotedByCurrentUser"`
		IsFollowing             bool       `json:"isFollowing"`
		Responses               []Response `json:"responses"`
		ID                      string     `json:"_id"`
		IsRepublished           bool       `json:"isRepublished"`
		FollowersCount          int        `json:"followersCount"`
		Author                  struct {
			BeingFollowed        bool          `json:"beingFollowed"`
			ID                   string        `json:"_id"`
			Role                 interface{}   `json:"role"`
			Name                 string        `json:"name"`
			Tagline              string        `json:"tagline"`
			Photo                string        `json:"photo"`
			Username             string        `json:"username"`
			CoverImage           string        `json:"coverImage"`
			NumReactions         int           `json:"numReactions"`
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
				Linkedin      string `json:"linkedin"`
				Stackoverflow string `json:"stackoverflow"`
				Google        string `json:"google"`
				Facebook      string `json:"facebook"`
				Twitter       string `json:"twitter"`
				Github        string `json:"github"`
				Website       string `json:"website"`
			} `json:"socialMedia"`
			StoriesCreated          []string `json:"storiesCreated"`
			NumFollowing            int      `json:"numFollowing"`
			NumFollowers            int      `json:"numFollowers"`
			IsDeactivated           bool     `json:"isDeactivated"`
			Location                string   `json:"location"`
			TotalAppreciationBadges int      `json:"totalAppreciationBadges"`
		} `json:"author"`
		Cuid               string    `json:"cuid"`
		Slug               string    `json:"slug"`
		Title              string    `json:"title"`
		Type               string    `json:"type"`
		V                  int       `json:"__v"`
		DateUpdated        time.Time `json:"dateUpdated"`
		ReactionToCountMap struct {
			Reaction5C090D96C2A9C2A674D35486 int `json:"reaction_5c090d96c2a9c2a674d35486"`
			Reaction5C090D96C2A9C2A674D35485 int `json:"reaction_5c090d96c2a9c2a674d35485"`
			Reaction5C090D96C2A9C2A674D35484 int `json:"reaction_5c090d96c2a9c2a674d35484"`
		} `json:"reactionToCountMap"`
		OgImage                string        `json:"ogImage"`
		ReactionsByCurrentUser []interface{} `json:"reactionsByCurrentUser"`
		TotalReactions         int           `json:"totalReactions"`
		Reactions              []string      `json:"reactions"`
		BookmarkedIn           []interface{} `json:"bookmarkedIn"`
		HasReward              bool          `json:"hasReward"`
		IsPublication          bool          `json:"isPublication"`
		NumCollapsed           int           `json:"numCollapsed"`
		DuplicatePosts         []interface{} `json:"duplicatePosts"`
		IsDelisted             bool          `json:"isDelisted"`
		AnsweredByTarget       bool          `json:"answeredByTarget"`
		Contributors           []struct {
			User struct {
				ID                   string        `json:"_id"`
				Username             string        `json:"username"`
				Name                 string        `json:"name"`
				Photo                string        `json:"photo"`
				Tagline              string        `json:"tagline"`
				Role                 interface{}   `json:"role"`
				CoverImage           string        `json:"coverImage"`
				NumReactions         int           `json:"numReactions"`
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
					Linkedin      string `json:"linkedin"`
					Stackoverflow string `json:"stackoverflow"`
					Google        string `json:"google"`
					Facebook      string `json:"facebook"`
					Twitter       string `json:"twitter"`
					Github        string `json:"github"`
					Website       string `json:"website"`
				} `json:"socialMedia"`
				StoriesCreated          []string `json:"storiesCreated"`
				NumFollowing            int      `json:"numFollowing"`
				NumFollowers            int      `json:"numFollowers"`
				IsDeactivated           bool     `json:"isDeactivated"`
				Location                string   `json:"location"`
				TotalAppreciationBadges int      `json:"totalAppreciationBadges"`
			} `json:"user"`
			Stamp string `json:"stamp"`
			ID    string `json:"_id"`
		} `json:"contributors"`
		IsEngaging      bool          `json:"isEngaging"`
		IsFeatured      bool          `json:"isFeatured"`
		IsActive        bool          `json:"isActive"`
		Followers       []interface{} `json:"followers"`
		ResponseCount   int           `json:"responseCount"`
		QuestionReplies []interface{} `json:"questionReplies"`
		DateAdded       time.Time     `json:"dateAdded"`
		UntaggedFrom    []interface{} `json:"untaggedFrom"`
		Tags            []struct {
			ID         string      `json:"_id"`
			Name       string      `json:"name"`
			Slug       string      `json:"slug"`
			IsApproved bool        `json:"isApproved"`
			IsActive   bool        `json:"isActive"`
			NumPosts   int         `json:"numPosts"`
			MergedWith interface{} `json:"mergedWith"`
			Logo       string      `json:"logo,omitempty"`
		} `json:"tags"`
		Downvotes      int `json:"downvotes"`
		Upvotes        int `json:"upvotes"`
		TotalPollVotes int `json:"totalPollVotes"`
		Reward         struct {
			Type string `json:"type"`
		} `json:"reward"`
		PollOptions     []interface{} `json:"pollOptions"`
		HasPolls        bool          `json:"hasPolls"`
		ContentMarkdown string        `json:"contentMarkdown"`
		Content         string        `json:"content"`
		Brief           string        `json:"brief"`
		CoverImage      string        `json:"coverImage"`
		Views           int           `json:"views"`
		IsAnonymous     bool          `json:"isAnonymous"`
	} `json:"post"`
}
type Response struct {
	ID              string `json:"_id"`
	Content         string `json:"content"`
	ContentMarkdown string `json:"contentMarkdown"`
	Author          struct {
		ID                   string        `json:"_id"`
		Username             string        `json:"username"`
		Name                 string        `json:"name"`
		Photo                string        `json:"photo"`
		Tagline              string        `json:"tagline"`
		Role                 interface{}   `json:"role"`
		CoverImage           string        `json:"coverImage"`
		NumReactions         int           `json:"numReactions"`
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
			Linkedin      string `json:"linkedin"`
			Stackoverflow string `json:"stackoverflow"`
			Google        string `json:"google"`
			Facebook      string `json:"facebook"`
			Twitter       string `json:"twitter"`
			Github        string `json:"github"`
			Website       string `json:"website"`
		} `json:"socialMedia"`
		StoriesCreated          []string `json:"storiesCreated"`
		NumFollowing            int      `json:"numFollowing"`
		NumFollowers            int      `json:"numFollowers"`
		IsDeactivated           bool     `json:"isDeactivated"`
		Location                string   `json:"location"`
		TotalAppreciationBadges int      `json:"totalAppreciationBadges"`
	} `json:"author"`
	Stamp              string `json:"stamp"`
	Post               string `json:"post"`
	V                  int    `json:"__v"`
	ReactionToCountMap struct {
		Reaction567453D0B73D6A82Ac8C5Abd int `json:"reaction_567453d0b73d6a82ac8c5abd"`
		Reaction5C090D96C2A9C2A674D35487 int `json:"reaction_5c090d96c2a9c2a674d35487"`
		Reaction5C090D96C2A9C2A674D35486 int `json:"reaction_5c090d96c2a9c2a674d35486"`
	} `json:"reactionToCountMap"`
	ReactionsByCurrentUser []interface{} `json:"reactionsByCurrentUser"`
	TotalReactions         int           `json:"totalReactions"`
	Reactions              []string      `json:"reactions"`
	Score                  int           `json:"score"`
	BookmarkedIn           []interface{} `json:"bookmarkedIn"`
	IsRewardWinner         bool          `json:"isRewardWinner"`
	TotalBadgesAwarded     int           `json:"totalBadgesAwarded"`
	BadgesAwarded          []interface{} `json:"badgesAwarded"`
	IsCollapsed            bool          `json:"isCollapsed"`
	Downvotes              int           `json:"downvotes"`
	Upvotes                int           `json:"upvotes"`
	DownvotedBy            []interface{} `json:"downvotedBy"`
	UpvotedBy              []interface{} `json:"upvotedBy"`
	IsActive               bool          `json:"isActive"`
	DateAdded              time.Time     `json:"dateAdded"`
	Popularity             float64       `json:"popularity"`
	Replies                []Reply       `json:"replies"`
}

type Reply struct {
	ReactionToCountMap struct {
		Any int `json:"any"`
	} `json:"reactionToCountMap"`
	Content         string `json:"content"`
	ContentMarkdown string `json:"contentMarkdown"`
	Author          struct {
		ID                   string        `json:"_id"`
		Username             string        `json:"username"`
		Name                 string        `json:"name"`
		Photo                string        `json:"photo"`
		Tagline              string        `json:"tagline"`
		CoverImage           string        `json:"coverImage"`
		Role                 string        `json:"role"`
		NumReactions         int           `json:"numReactions"`
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
			Linkedin      string `json:"linkedin"`
			Stackoverflow string `json:"stackoverflow"`
			Google        string `json:"google"`
			Facebook      string `json:"facebook"`
			Twitter       string `json:"twitter"`
			Github        string `json:"github"`
			Website       string `json:"website"`
		} `json:"socialMedia"`
		StoriesCreated          []string `json:"storiesCreated"`
		NumFollowing            int      `json:"numFollowing"`
		NumFollowers            int      `json:"numFollowers"`
		IsDeactivated           bool     `json:"isDeactivated"`
		Location                string   `json:"location"`
		TotalAppreciationBadges int      `json:"totalAppreciationBadges"`
	} `json:"author"`
	Stamp                  string        `json:"stamp"`
	ID                     string        `json:"_id"`
	ReactionsByCurrentUser []interface{} `json:"reactionsByCurrentUser"`
	TotalReactions         int           `json:"totalReactions"`
	Reactions              []interface{} `json:"reactions"`
	Upvotes                int           `json:"upvotes"`
	IsActive               bool          `json:"isActive"`
	DateAdded              time.Time     `json:"dateAdded"`
}
