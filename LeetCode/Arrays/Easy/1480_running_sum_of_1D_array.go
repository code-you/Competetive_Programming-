package main

import "fmt"

//O(n) time and space complexity
func runningSum(nums []int) []int {
	ans := make([]int, len(nums))
	ans[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] + nums[i]
	}
	return ans
}

func main() {
	arr := []int{1, 2, 3, 4}
	nums := runningSum(arr)
	fmt.Println(nums)
}
