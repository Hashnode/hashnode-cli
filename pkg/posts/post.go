package posts

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
	blackfriday "github.com/russross/blackfriday"
	"github.com/scriptonist/termd/pkg/console"
)

// openPost opens a post in a new tview box
func openPost(app *tview.Application, postcuid string, list *tview.List) {

	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true).
		SetWordWrap(true).
		SetTextAlign(tview.AlignLeft).
		SetChangedFunc(func() {
			app.Draw()
		})

	textView.Box = textView.Box.SetBorder(true).SetBorderPadding(1, 1, 2, 1)
	textView.SetBorder(true)

	go func() {
		if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
			app.Stop()
			panic(err)
		}
	}()

	var singlePost Post
	textView.Write([]byte("[:green:l]Loading....[-:-:-]"))
	b, err := makeRequest(fmt.Sprintf("%s/%s", postAPI, postcuid))
	if err != nil {
		app.Stop()
		log.Fatal(err)
	}
	textView.Clear()

	err = json.Unmarshal(b, &singlePost)
	if err != nil {
		app.Stop()
		log.Fatal(err)
	}

	title := fmt.Sprintf("\nTitle: %s", singlePost.Post.Title)
	var author string
	if singlePost.Post.Author.Name != "" {
		author = fmt.Sprintf("Author: %s", singlePost.Post.Author.Name)
	} else {
		author = fmt.Sprintf("Author: Anonymous")
	}

	reactions := fmt.Sprintf("Reactions: %d", singlePost.Post.TotalReactions)
	ptype := fmt.Sprintf("Type: %s", singlePost.Post.Type)
	link := fmt.Sprintf("Link: https://hashnode.com/post/%s", singlePost.Post.Cuid)
	writeToTextView(textView,
		title,
		author,
		reactions,
		ptype,
		link,
		"\n",
		renderTerminal(singlePost.Post.ContentMarkdown),
		func() string {
			if len(singlePost.Post.Responses) > 0 {
				return fmt.Sprintf("\n%s\n%s\n", "[green]Responses[white]",
					"[green]==========[white]")
			}
			return ""
		}(),
	)
	for ind, response := range singlePost.Post.Responses {
		writeToTextView(
			textView,
			fmt.Sprintf("\n%d", ind+1),
			fmt.Sprintf("---"),
			renderTerminal(response.ContentMarkdown),
		)
		if len(response.Replies) > 0 {
			writeToTextView(textView,
				"\n\t[green]Replies[white]",
				"\t[green]=======[white]",
			)
			for indreply, reply := range response.Replies {
				writeToTextView(
					textView,
					fmt.Sprintf("\n\t%d", indreply+1),
					fmt.Sprintf("\t---"),
					fmt.Sprintf("\tAuthor: %s", reply.Author.Name),
					fmt.Sprintf("\t%s", renderTerminal(reply.ContentMarkdown)),
				)

			}
		}

	}
	textView.ScrollToBeginning()

	textView.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
				app.Stop()
				panic(err)
			}
		}
	})

}

func writeToTextView(t *tview.TextView, contents ...string) {
	for _, content := range contents {
		t.Write([]byte(content))
		t.Write([]byte("\n"))
	}
}

func renderTerminal(content string) string {
	r := console.Console{}
	out := string(blackfriday.Run([]byte(content),
		blackfriday.WithRenderer(r),
		blackfriday.WithExtensions(blackfriday.CommonExtensions)))
	return out

}
