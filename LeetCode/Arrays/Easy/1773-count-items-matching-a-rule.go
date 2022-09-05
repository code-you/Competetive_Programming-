//Ques:- https://leetcode.com/problems/count-items-matching-a-rule/
package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	id := -1
	count := 0
	if ruleKey == "type" {
		id = 0
	} else if ruleKey == "color" {
		id = 1
	} else {
		id = 2
	}
	for _, v := range items {
		if v[id] == ruleValue {
			count++
		}

	}

	return count
}

//func main() {
//	arr := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "iphone"}}
//	nums := countMatches(arr, "color", "silver")
//	fmt.Println(nums)
//}
