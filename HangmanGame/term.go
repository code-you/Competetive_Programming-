package main

import "fmt"

type HangmanTerm struct {
}

func (h *HangmanTerm) RenderGame(placeholder []string, entries map[string]bool) error {
	fmt.Println()
	fmt.Println(placeholder)
	fmt.Printf("Chances: %d\n", MaxChances-len(entries))
	fmt.Printf("Entries: %v\n", get_Keys(entries))
	fmt.Println("Guess a letter or the word :")

	return nil
}
func (h *HangmanTerm) GetInput() string {
	str := ""
	fmt.Scanln(&str)

	return str
}
