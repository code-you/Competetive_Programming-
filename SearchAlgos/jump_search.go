package main

import (
	"fmt"
	"math"
)

// It only for sorted arrays     ---->linearSearch<JumpSearch<BinarySearch<---- Increasing order of searching
func jumpSearch(list []int, key int) int {
	n := len(list)
	// Finding block size to be jumped
	step := math.Sqrt(float64(n))

	// Finding the block where element is
	// present (if it is present)
	prev := 0
	for list[int(math.Min(step, float64(n)))-1] < key {
		prev = int(step)
		step += math.Sqrt(float64(n))
		if prev >= n {
			return -1
		}

	}

	// Doing a linear search for x in block
	// beginning with prev.
	for list[prev] < key {
		prev++

		// If we reached next block or end of
		// array, element is not present.
		if prev == int(math.Min(step, float64(n))) {
			return -1
		}

	}
	// If element is found
	if list[prev] == key {

		return prev
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("index:", jumpSearch(arr, 5))
}
