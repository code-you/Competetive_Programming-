// Ques:- https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	X := 0
	for _, v := range operations {
		if v == "++X" {
			X += 1
		} else if v == "--X" {
			X -= 1
		} else if v == "X++" {
			X++
		} else {
			X--
		}

		// OR

		// if strings.Contains(v,"+"){
		//     X++
		// }else{
		//     X--
		// }
	}
	return X
}

func main() {
	arr := []string{"--X", "X++", "X++"}
	nums := finalValueAfterOperations(arr)
	fmt.Println(nums)
}
