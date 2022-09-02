package main

import "fmt"

func removeDuplicates(nums []int) int {

	count := make(map[int]int)
	i := 0

	for _, value := range nums {
		if count[value] < 2 {
			count[value] = count[value] + 1
			nums[i] = value
			i++
		}
	}

	return i
}

func main() {

	nums := []int{1, 1, 1, 2, 2, 3}

	n := removeDuplicates(nums)
	fmt.Println(n)
}
