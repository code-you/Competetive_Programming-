package main

import "fmt"

func LinearSearch(list []int, key int) bool {
	for _, val := range list {
		if val == key {
			return true
		}
	}
	return false
}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 14, 12, 11}
	fmt.Println(LinearSearch(arr, 8))  // true
	fmt.Println(LinearSearch(arr, 10)) // false

}
