package posts

import (
	"encoding/json"
	"fmt"
	"log"
)

func GetNews() {
	b, err := makeRequest(newsAPI)
	if err != nil {
		log.Println(err)
	}
	m := make(map[string]interface{})
	json.Unmarshal(b, &m)
	fmt.Println(m)
}
