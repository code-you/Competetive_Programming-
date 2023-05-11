//Ques:- https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
package main

import (
	"fmt"
)

func average(salary []int) float64 {

	sum := 0
	//sort.Ints(salary)
	//
	//for i := 1; i < len(salary)-1; i++ {
	//	sum += salary[i]
	//}
	////ratio := math.Pow(10, 5)
	//
	//return float64(sum) / float64(len(salary)-2)

	minSalary, maxSalary := salary[0], salary[0]
	for _, v := range salary {
		sum += v
		if v > maxSalary {
			maxSalary = v
		}
		if v < minSalary {
			minSalary = v
		}
	}

	return float64(sum-maxSalary-minSalary) / float64(len(salary)-2)
}

func main() {
	arr := []int{48000, 59000, 99000, 13000, 78000, 45000, 31000, 17000, 39000, 37000, 93000, 77000, 33000, 28000, 4000, 54000, 67000, 6000, 1000, 11000}
	val := average(arr)
	fmt.Println(val)
}
