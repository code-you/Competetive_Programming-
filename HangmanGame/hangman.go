package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var MaxChances = 8
var dev = flag.Bool("dev", false, "dev mode")

func get_Keys(entries map[string]bool) (keys []string) {
	for k, _ := range entries {
		keys = append(keys, k)
	}

	return
}
func get_word() string {
	if *dev {
		return "elephant"
	}
	resp, err := http.Get("https://random-word-api.herokuapp.com/word")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var words []string
	err = json.Unmarshal(body, &words)
	if err != nil {
		return "elephant"
	}
	for _, word := range words {
		if len(word) > 4 && len(word) < 9 {
			return word
		}
	}
	return words[0]
}

func main() {
	flag.Parse()
	word := get_word()
	h := HangmanTerm{}
	result := Play(&h, word)
	if result {
		fmt.Println("You Win!")
	} else {
		fmt.Println("You loss , word was:", word)

	}
}
func Play(h *HangmanTerm, word string) bool {

	//lookup for entries

	entries := map[string]bool{}

	//list of "_" corresponding to the letter
	var placeholder []string

	for i := 0; i < len(word); i++ {
		placeholder = append(placeholder, "_")
	}

	t := time.NewTimer(5 * time.Second)

	chances := MaxChances
	result := make(chan bool)
	go func() {
		for {
			userInput := strings.Join(placeholder, "")
			if chances == 0 && userInput != word {
				result <- false

				return
			}
			if userInput == word {
				result <- true

				return
			}
			h.RenderGame(placeholder, entries)
			str := h.GetInput()
			if len(str) > 2 {
				if str == word {
					result <- true

					return
				} else {
					entries[str] = true
					chances -= 1
					continue
				}
			}
			_, ok := entries[str]
			if ok {
				continue
			}

			entries[str] = true
			found := false
			for i, v := range word {
				if str == string(v) {
					placeholder[i] = string(v)
					found = true
				}
			}
			if !found {
				chances -= 1
			}
		}
	}()

	select {
	case r := <-result:
		if r {
			return true
		} else {
			return false
		}
		//goto END
	case <-t.C:
		return false
		//goto END

	}

}
