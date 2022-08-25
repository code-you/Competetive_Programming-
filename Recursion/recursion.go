package main

import "fmt"

var count = 0

func printVal() {
	if count == 2 { // Base Case
		return
	}
	fmt.Println(count)
	count++
	printVal()
}

func printName(i, n int) {
	if i > n {
		return
	}
	fmt.Println("Aditya")

	printName(i+1, n)
}

func printFromOneToN(i, n int) {
	if i > n {
		return
	}
	fmt.Println(i)

	printFromOneToN(i+1, n)
}

func printFromNToOne(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printFromNToOne(n - 1)
}

func printFromOneToNUsingBackTracing(n int) {
	if n == 0 {
		return
	}
	printFromOneToNUsingBackTracing(n - 1)
	fmt.Println(n)
}

func printFromNToOneUsingBackTracing(i, n int) {
	if i > n {
		return
	}
	printFromNToOneUsingBackTracing(i+1, n)
	fmt.Println(i)
}

func SumOfN(i, sum int) { // using parameterized way
	if i < 1 {
		print(sum)
		return
	}
	SumOfN(i-1, sum+i)
}

func SumOfN2(n int) int { // using functions way
	if n == 0 {
		return 0
	}

	return n + SumOfN2(n-1)
}

func Factorial(n int) int { // using functions way
	if n == 0 {
		return 1
	}

	return n * Factorial(n-1)
}

func ReverseArray(nums []int, i, n int) {
	if i >= n/2 {
		return
	}
	nums[n-i-1], nums[i] = nums[i], nums[n-i-1] // swapping to variales
	ReverseArray(nums, i+1, n)
}

func Palindrome(i int, s string, n int) bool {
	if i >= n/2 {
		return true
	}
	if s[i] != s[n-i-1] {
		return false
	}
	return Palindrome(i+1, s, n)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	var n int
	fmt.Scan(&n)
	printVal()
	printName(1, n)
	printFromOneToN(1, n)
	printFromNToOne(5)
	printFromOneToNUsingBackTracing(n)
	printFromNToOneUsingBackTracing(1, n)
	SumOfN(n, 0)
	fmt.Println(SumOfN2(n))
	fmt.Println(Factorial(n))

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	ReverseArray(arr, 0, len(arr))
	fmt.Println(arr)

	s := "MADAM"
	fmt.Println(Palindrome(0, s, len(s)))

	fmt.Println(Fibonacci(n))
}
