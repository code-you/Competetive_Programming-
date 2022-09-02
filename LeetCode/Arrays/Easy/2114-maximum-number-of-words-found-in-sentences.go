//Ques:- https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
package main

import "strings"

func mostWordsFound(sentences []string) int {
	x := 0
	for _, str := range sentences {
		val := strings.Split(str, " ")
		if x < len(val) {
			x = len(val)
		}
	}
	return x
}
