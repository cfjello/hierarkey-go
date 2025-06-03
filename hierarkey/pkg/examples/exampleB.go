package examples

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/yourusername/hierarkey-go/hierarkey"
)

func RunExampleB() {
	hk := hierarkey.New(1, 2)

	resp, err := http.Get("https://en.wikipedia.org/w/api.php?action=parse&page=List_of_popular_music_genres&section=15&prop=wikitext&format=json")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var electronic map[string]interface{}
	if err := json.Unmarshal(body, &electronic); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	parseData := electronic["parse"].(map[string]interface{})
	wikitextData := parseData["wikitext"].(map[string]interface{})
	content := wikitextData["*"].(string)

	entries := regexp.MustCompile(`\r?\n`).Split(content, -1)
	level := 0
	prevLevel := 0

	for _, item := range entries {
		show := true
		if strings.HasPrefix(item, "==") {
			hk.JumpToLevel("0")
			level = 0
			prevLevel = 0
		} else if strings.HasPrefix(item, "*") {
			// Count the number of asterisks to determine level
			level = len(strings.TrimLeft(item, "*")) - len(strings.TrimLeft(item, " *"))

			if level > prevLevel {
				hk.NextLevel()
			} else if level == prevLevel {
				hk.NextLeaf()
			} else {
				hk.PrevLevel(prevLevel - level)
			}
			prevLevel = level
		} else {
			// Ignore all other lines not related directly to the music genre entries
			show = false
		}

		if show {
			entry := regexp.MustCompile(`[=\*\[\]]`).ReplaceAllString(item, "")
			fmt.Printf("%s: %s\n", hk.CurrLeaf(), entry)
		}
	}
}
