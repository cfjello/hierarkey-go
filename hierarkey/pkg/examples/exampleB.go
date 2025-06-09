package examples

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/cfjello/hierarkey-go/hierarkey/pkg/hierarkey/pkg/hierarkey"
)

func RunExampleB() {
	hk := hierarkey.NewHierarKey(1, 2)

	resp, err := http.Get("https://en.wikipedia.org/w/api.php?action=parse&page=List_of_music_genres_and_styles&section=13&prop=wikitext&format=json")
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

	var rockMap map[string]any
	if err := json.Unmarshal(body, &rockMap); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	parseData, ok := rockMap["parse"].(map[string]any)
	if !ok {
		fmt.Println("Error: parse data not found or has unexpected format")
		return
	}
	wikitextData := parseData["wikitext"].(map[string]any)
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
			level = len(item) - len(strings.TrimLeft(item, "*"))

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
			fmt.Printf("%s: %s\n", hk.GetCurrLeaf(), entry)
		}
	}
}
