//Ques :- https://leetcode.com/problems/richest-customer-wealth/

package main

func maximumWealth(accounts [][]int) int {
	max := 0
	for i := range accounts {
		val := sumOfArray(accounts[i])
		if val >= max {
			max = val
		}
	}
	return max
}

func sumOfArray(nums []int) int {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	return sum
}

//func main() {
//	arr := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
//	nums := maximumWealth(arr)
//	fmt.Println(nums)
//}
