// Ques:- https://leetcode.com/problems/build-array-from-permutation/

package main

import "fmt"

func buildArray(nums []int) []int {
	//space complexity O(1)
	n := len(nums)
	for i := range nums {
		nums[i] = n*(nums[nums[i]]%n) + nums[i]
	}
	for i := range nums {
		nums[i] /= n
	}
	return nums

	// space complexity O(n)
	//ans := make([]int, len(nums))
	//for i := range nums {
	//	ans[i] = nums[nums[i]]
	//}
	//return ans
}

func main() {
	arr := []int{0, 2, 1, 5, 3, 4}
	nums := buildArray(arr)
	fmt.Println(nums)
}
