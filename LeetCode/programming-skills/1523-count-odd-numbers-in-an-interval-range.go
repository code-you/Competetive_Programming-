//Ques:- https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/
package main

func countOdds(low int, high int) int {
	//count := 0
	//for low <= high {
	//	if low%2 != 0 {
	//		count++
	//
	//	}
	//	low++
	//}
	//return count
	if low%2 == 0 && high%2 == 0 {
		return (high - low) / 2
	} else {
		return (high-low)/2 + 1
	}
}
