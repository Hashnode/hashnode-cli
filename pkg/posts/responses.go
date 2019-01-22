package posts

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/rivo/tview"
)

func openResponses(t *tview.TextView, postcuid string, totalResponses int) {

	const defaultSortOrder = "totalReactions"

	r, err := getResponses(postcuid, totalResponses, 1, defaultSortOrder)
	if err != nil {
		log.Println(err)
	}
	var responsesAPI responsesAPI
	err = json.Unmarshal(r, &responsesAPI)
	if err != nil {
		log.Println(err)
	}

	noresponse := len(responsesAPI.Responses)
	for ind, response := range responsesAPI.Responses {
		writeToTextView(
			t,
			fmt.Sprintf("\n[green]Response %d/%d[white]", ind+1, noresponse),
			fmt.Sprintf("[green]--------------[green]"),
			renderTerminal(response.ContentMarkdown),
		)
		if len(response.Replies) > 0 {
			writeToTextView(t,
				"\n[yellow]Replies[white]",
				"[yellow]=======[white]",
			)
			noreplies := len(response.Replies)
			for indreply, reply := range response.Replies {
				writeToTextView(
					t,
					fmt.Sprintf("\n[yellow]Reply %d/%d[white]", indreply+1, noreplies),
					fmt.Sprintf("[yellow]~~~~~~~~~~~[white]"),
					fmt.Sprintf("Author: %s", reply.Author.Name),
					indentMarkdown(renderTerminal(reply.ContentMarkdown), "\t"),
				)

			}
		}

	}

}

func getResponses(postID string, perPage, page int, sortOrder string) ([]byte, error) {
	const apiURL = "https://hashnode.com/ajax/responses"
	u, err := url.Parse(apiURL)
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Set("post_id", postID)
	q.Set("page", fmt.Sprintf("%d", page))
	q.Set("per_page", fmt.Sprintf("%d", perPage))
	q.Set("sort_order", sortOrder)
	u.RawQuery = q.Encode()

	client := getHttpClient()
	resp, err := client.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, err
}

type responsesAPI struct {
	Pagination struct {
		Page    string `json:"page"`
		PerPage string `json:"per_page"`
		Total   int    `json:"total"`
	} `json:"pagination"`
	Order     string `json:"order"`
	Responses []Response
}
