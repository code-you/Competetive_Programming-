package main

import "fmt"

func Search(list []int, key int) bool {
	start := 0
	end := len(list) - 1
	for start <= end {
		mid := start + (end-start)/2
		if list[mid] == key {
			return true
		}
		if list[mid] > key {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(Search(arr, 1))  // true
	fmt.Println(Search(arr, 10)) // false

}
