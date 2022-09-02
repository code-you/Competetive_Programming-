//Ques:- https://leetcode.com/problems/number-of-good-pairs/
package main

func numIdenticalPairs(nums []int) int {

	// Brute-Force approach
	count := 0
	//for i := 0; i < len(nums); i++ {
	//	for j := i + 1; j < len(nums); j++ {
	//		if nums[i] == nums[j] && i < j {
	//			count++
	//		}
	//	}
	//}
	//return count

	// optimised approach we can take an array length just greater than our constraints as in
	//our case constraints is 100, so we take just one greater size array
	temp := make([]int, 101)
	for i := 0; i < len(nums); i++ {
		count += temp[nums[i]]
		temp[nums[i]]++
	}

	return count
}
