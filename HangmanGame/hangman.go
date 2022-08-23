package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func get_Keys(entries map[string]bool) (keys []string) {
	for k, _ := range entries {
		keys = append(keys, k)
	}

	return
}
func get_word() string {
	resp, err := http.Get("https://random-word-api.herokuapp.com/word")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var word []string
	err = json.Unmarshal(body, &word)
	if err != nil {
		return "elephant"
	}

	return word[0]
}

func main() {
	fmt.Println(get_word())
}
