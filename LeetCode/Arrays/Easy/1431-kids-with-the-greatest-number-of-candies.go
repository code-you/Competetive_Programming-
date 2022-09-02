//Ques:- https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
package main

func kidsWithCandies(candies []int, extraCandies int) []bool {

	res := make([]bool, len(candies))
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies > i {
			res[i] = true
		} else {
			res[i] = false
		}
	}

	return res
}
