package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | float64 | float32
}

func BubbleSort[N Number](input []N) []N {
	n := len(input)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if input[i] > input[i+1] {
				input[i], input[i+1] = input[i+1], input[i]
				swapped = true
			}
		}
	}
	return input
}

func main() {
	arr := []int{1, 0, -1, 4, 3, 2, -4}
	fmt.Println(BubbleSort(arr))
}
