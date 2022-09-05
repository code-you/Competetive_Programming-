//Ques :- https://leetcode.com/problems/concatenation-of-array/

package main

//O(n) time and space complexity
func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)
	for i := range nums {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}

	return ans
}

//func main() {
//	arr := []int{1, 3, 2, 1}
//	nums := getConcatenation(arr)
//	fmt.Println(nums)
//}
