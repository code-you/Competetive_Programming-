//Ques:- https://leetcode.com/problems/shuffle-string/
package main

func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i, v := range indices {
		res[v] = s[i]
	}
	return string(res)
}
