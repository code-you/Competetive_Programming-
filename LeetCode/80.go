package main

import (
	"strings"
)

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

	nums := []int{194, 155, 231, 184, 185, 246, 176, 131, 161, 222, 174, 227, 162, 134, 241, 154, 168, 185, 218, 178, 229, 187, 139, 246, 178, 187, 139, 204, 146, 225, 148, 179, 245, 139, 172, 134, 193, 156, 233, 131, 154, 240, 166, 188, 190, 216, 150, 230, 145, 144, 240, 167, 140, 163, 221, 190, 238, 168, 139, 241, 154, 159, 164, 199, 170, 224, 173, 140, 244, 182, 143, 134, 206, 181, 227, 172, 141, 241, 146, 159, 170, 202, 134, 230, 142, 163, 244, 172, 140, 191}

	//n := removeDuplicates(nums)
	//n := validUtf8(nums)
	//fmt.Println(n)
}
func mergeAlternately(word1 string, word2 string) string {
	s := strings.Builder{}
	i := 0

	// Allocate enough space to merged string
	s.Grow(len(word1) + len(word2))

	// Loop through word1 and word 2
	for i < len(word1) || i < len(word2) {

		// Add letter from word1 if it exists
		if i < len(word1) {
			s.WriteByte(word1[i])
		}

		// Add letter from word2 if it exists
		if i < len(word2) {
			s.WriteByte(word2[i])
		}

		i++
	}

	// Return merged
	return s.String()
}

//func fizzBuzz(n int) []string {
//	m := make([]string, n)
//	result := ""
//	for i := 1; i < n; i++ {
//		if i%3 == 0 && i%5 == 0 {
//			result += "FizzBuzz"
//		} else if i%3 == 0 {
//			result += "Fizz"
//		} else if i%5 == 0 {
//			result += "Buzz"
//		} else {
//			result += strconv.Itoa(i)
//		}
//	}
//	m = append(m, result)
//	return m
//}
//
//func numberOfSteps(num int) int {
//	count := 0
//	for num > 0 {
//		if num%2 == 0 {
//			num /= 2
//
//		} else {
//			num--
//		}
//		count++
//	}
//	return count
//}
//
//func hammingWeight(num uint32) int {
//	strconv.Itoa(int(num))
//	count := 0
//	for num > 0 {
//		rem := num % 10
//		if rem == 1 {
//			count++
//		}
//		num /= 10
//	}
//	return count
//}
//
//func reverseBits(num uint32) uint32 {
//	rev := uint32(0)
//
//	// traversing bits of 'n'
//	// from the right
//	for num > 0 {
//		// bitwise left shift
//		// 'rev' by 1
//		rev <<= 1
//
//		// if current bit is '1'
//		if (num & 1) == 1 {
//			rev ^= 1
//		}
//
//		// bitwise right shift
//		//'n' by 1
//		num >>= 1
//	}
//	// required number
//	return rev
//}
//
//func subtractProductAndSum(n int) int {
//	sum := 0
//	mul := 1
//	for n != 0 {
//		rem := n % 10
//		mul *= rem
//		sum += rem
//		n /= 10
//	}
//	return int(math.Abs(float64(mul - sum)))
//}
//
//func largestPerimeter(nums []int) int {
//	result := 0
//	sort.Ints(nums)
//	for i := 0; i < len(nums)-2; i++ {
//		if nums[i+2] < nums[i+1]+nums[i] {
//			result = nums[i+2] + nums[i+1] + nums[i]
//		}
//	}
//
//	return result
//}
//
//func arraySign(nums []int) int {
//	product := float64(1)
//	for _, v := range nums {
//		product *= float64(v)
//	}
//	if product > 0 {
//		return 1
//	} else if product < 0 {
//		return -1
//	}
//
//	return 0
//}
//
//func canMakeArithmeticProgression(arr []int) bool {
//	sort.Ints(arr)
//	diff := arr[1] - arr[0]
//	for i := 1; i < len(arr)-1; i++ {
//		val := arr[i+1] - arr[i]
//		if val == diff {
//			continue
//		} else {
//			return false
//		}
//
//	}
//	return true
//}
//
//func isHappy(n int) bool {
//	if math.Floor(math.Log10(float64(n)))+1 == 1 {
//		return false
//	}
//	for {
//		n = sumOfDigits(n)
//		if sum != 1 {
//			continue
//		} else {
//			return true
//		}
//
//	}
//}
//
//func sumOfDigits(val int) int {
//	sum := 0
//	for val != 0 {
//		rem := val % 10
//		sum += rem * rem
//		val /= 10
//	}
//	return sum
//}
//
//func numberOfWeakCharacters(properties [][]int) int {
//	count := 0
//	for i := 0; i < len(properties); i++ {
//		for j := 0; j < len(properties[i])-1; j++ {
//			if properties[i][j] < properties[i][j+1] {
//				count++
//			} else {
//				continue
//			}
//		}
//	}
//	return count
//}
//
//func maxProfit(k int, prices []int) int {
//	cur := [2][]int{}
//
//	for ind := len(prices) - 1; ind >= 0; ind-- {
//		for buy := 0; buy <= 1; buy++ {
//			for capa := 1; capa <= k; capa++ {
//				if buy == 1 {
//					cur[buy][capa] = math.Max()
//				}
//			}
//		}
//	}
//}
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func middleNode(head *ListNode) *ListNode {
//	middle := head
//	end := head
//
//	for end != nil && end.Next != nil {
//		middle = middle.Next
//		end = end.Next.Next
//	}
//
//	return middle
//}

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == 0 {
				return false
			}
			if arr[i] == 2*arr[j] {
				println(arr[i], arr[j])
				return true
			}
		}
	}
	return false
}

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	l := len(arr)
	for i := 0; i < l; i++ {
		sum += ((i+1)*(l-i) + 1) / 2 * arr[i]
	}

	return sum
}

func moveZeroes(nums []int) {
	m := make([]int, len(nums))

	for _, v := range nums {
		if v != 0 {
			m = append(m, v)
		}
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = m[i]
	}
}

func toLowerCase(s string) string {
	j := strings.ToLower(s)
	return j
}

func isAlienSorted(words []string, order string) bool {
	characters := map[uint8]int{}
	for i := 0; i < len(order); i++ {
		characters[order[i]] = i
	}

	prev := words[0]
	for i := 1; i < len(words); i++ {
		var j int
		for j = 0; j < len(prev)-1 && j < len(words[i])-1 && prev[j] == words[i][j]; j++ {
		}

		if characters[prev[j]] > characters[words[i][j]] || characters[prev[j]] == characters[words[i][j]] && len(prev) > len(words[i]) {
			return false
		}
		prev = words[i]
	}
	return true
}
