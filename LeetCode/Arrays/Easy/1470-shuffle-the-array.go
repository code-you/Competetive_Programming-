// Ques :- https://leetcode.com/problems/shuffle-the-array/
package main

func shuffle(nums []int, n int) []int {
	ans := make([]int, len(nums))

	//1st approach Brute force approach

	i, j, k := 0, 0, n
	for i < len(nums) && j < len(nums) && k < len(nums) {
		if i%2 == 0 {
			ans[i] = nums[j]
			j++
		} else {
			ans[i] = nums[k]
			k++
		}
		i++
	}

	//2nd Brute force approach

	//j := 0
	//for i := 0; i < 2*n; i += 2 {
	//	ans[i] = nums[j]
	//	ans[i+1] = nums[j+n]
	//	j++
	//}
	//return ans

	//Optimized approach using assignment of 2 values at one places 1*10000+3= 10003
	for i := 0; i < n; i++ {
		nums[i+n] += 10000 * nums[i] // on the place of 10000 we can use any value greater than 1000 because constrains
		// is from 1 to 1000, so we can take any value greater than 1000 like 1005 then replace from extraction also....
	}
	for i := 0; i < n; i++ {
		nums[2*i] = nums[n+i] / 10000 // for even places

		nums[2*i+1] = nums[n+i] % 10000 // for odd places
	}

	// Using Bitwise operator

	for i := 0; i < n; i++ {
		nums[i+n] = (nums[i+n] << 10) | nums[i] // left shift with 10 because constrains is 1000, and it can be store using 10 bits the number at nums[i+n] and the OR with nums[i]
	}
	for i := 0; i < n; i++ {
		nums[2*i] = nums[i+n] & 1023  // & AND by 1023 to get actual OR value
		nums[2*i+1] = nums[i+n] >> 10 // right shift to get the value
	}
	return nums

}
