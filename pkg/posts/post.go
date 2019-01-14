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
	}()

	modal := tview.NewModal().SetText("Loading....")
	go func() {
		if err := app.SetRoot(modal, false).SetFocus(modal).Run(); err != nil {
			app.Stop()
			panic(err)
		}
	}()

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
	noresponse := len(singlePost.Post.Responses)
	for ind, response := range singlePost.Post.Responses {
		writeToTextView(
			textView,
			fmt.Sprintf("\n[green]Response %d/%d[white]", ind+1, noresponse),
			fmt.Sprintf("[green]--------------[green]"),
			renderTerminal(response.ContentMarkdown),
		)
		if len(response.Replies) > 0 {
			writeToTextView(textView,
				"\n[yellow]Replies[white]",
				"[yellow]=======[white]",
			)
			noreplies := len(response.Replies)
			for indreply, reply := range response.Replies {
				writeToTextView(
					textView,
					fmt.Sprintf("\n[yellow]Reply %d/%d[white]", indreply+1, noreplies),
					fmt.Sprintf("[yellow]~~~~~~~~~~~[white]"),
					fmt.Sprintf("Author: %s", reply.Author.Name),
					indentMarkdown(renderTerminal(reply.ContentMarkdown), "\t"),
				)

			}
		}

	}

	textView.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
				app.Stop()
				panic(err)
			}
		}
	})

	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
		app.Stop()
		panic(err)
	}

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

func indentMarkdown(s string, prefix string) string {
	// var lines []string
	// for _, line := range strings.Split(s, "\n") {
	// 	lines = append(lines, fmt.Sprintf("%s%s", prefix, line))
	// }
	return fmt.Sprintf("%s%s", prefix, s)
}
