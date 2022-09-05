//Ques:- https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
package main

func kidsWithCandies(candies []int, extraCandies int) []bool {

	//1st-Brute-force approach
	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		test := false
		for j := 0; j < len(candies); j++ {
			if candies[i]+extraCandies < candies[j] {
				test = true
				break
			}
		}
		if test {
			res[i] = false
		} else {
			res[i] = true
		}
	}
	return res

	//2nd-optimised approach
	//res := make([]bool, len(candies))
	//max := 0
	//for _, v := range candies {
	//	if v >= max {
	//		max = v
	//	}
	//}
	//for i := 0; i < len(candies); i++ {
	//	if candies[i]+extraCandies >= max {
	//		res[i] = true
	//	} else {
	//		res[i] = false
	//	}
	//}
	//
	//return res
}
