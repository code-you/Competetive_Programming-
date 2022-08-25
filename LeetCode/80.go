//package main
//
//import "fmt"
//
//func removeDuplicates(nums []int) int {
//
//	count := make(map[int]int)
//	i := 0
//
//	for _, value := range nums {
//		if count[value] < 2 {
//			count[value] = count[value] + 1
//			nums[i] = value
//			i++
//		}
//	}
//
//	return i
//}
//
//func main() {
//
//	nums := []int{1, 1, 1, 2, 2, 3}
//
//	n := removeDuplicates(nums)
//	fmt.Println(n)
//}

package main

import (
	"fmt"
)

func main() {
	// your code goes here
	var T int
	fmt.Scan(&T)

	for T > 0 {
		var N, M int
		fmt.Scanln(&N, &M)
		if (N/M)%2 == 0 && M%2 == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

		T--
	}

}
