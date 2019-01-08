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

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	rootAPIURL = "https://hashnode.com/ajax"
)

// API URL's
var (
	hotPostsAPI = fmt.Sprintf("%s/posts/%s", rootAPIURL, "hot")
	newsAPI     = fmt.Sprintf("%s/posts/%s", rootAPIURL, "news")
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
	// box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	// if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
	// 	panic(err)
	// }

}

func openPost(app *tview.Application, postcuid string, list *tview.List) {
	var singlePost Post
	b, err := makeRequest(fmt.Sprintf("%s/%s", postAPI, postcuid))
	if err != nil {
		app.Stop()
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &singlePost)
	if err != nil {
		app.Stop()
		log.Fatal(err)
	}

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	textView.SetText(singlePost.Post.ContentMarkdown)

	textView.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
				app.Stop()
				panic(err)
			}
		}
	})
	textView.SetBorder(true)
	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
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
