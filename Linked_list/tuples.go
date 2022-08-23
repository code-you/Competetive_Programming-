package main

import "fmt"

//gets the power series of integer a and returns tuple of square of a
// and cube of a
func powerSeries(a int) (int, int) {
	return a * a, a * a * a
}

// main method
func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)
	fmt.Println("Square ", square, "Cube", cube)
}
