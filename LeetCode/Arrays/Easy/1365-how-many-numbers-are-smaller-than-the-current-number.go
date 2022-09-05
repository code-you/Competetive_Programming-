//Ques:- https://leetcode.com/problems/how-many-numbers-are-smaller-than-the-current-number/

package main

func smallerNumbersThanCurrent(nums []int) []int {

	//1- Brute-force approach
	//res := make([]int, len(nums))
	//for i := 0; i < len(nums); i++ {
	//	count := 0
	//	for j := 0; j < len(nums); j++ {
	//		if nums[i] > nums[j] {
	//			count++
	//		}
	//	}
	//	res[i] = count
	//}
	//return res

	//2- Optimised approach
	res := make([]int, len(nums))
	temp := make([]int, 101) // why len 101 because constraints is 100

	//frequency array is ready how many times which element is occurred
	for i := 0; i < len(nums); i++ {
		temp[nums[i]]++
	}

	//running sum of the frequency array
	for i := 1; i < 101; i++ {
		temp[i] += temp[i-1]
	}

	// making it to the result
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			res[i] = 0
			continue
		}
		res[i] = temp[nums[i]-1]
	}

	return res
}
