//Ques:- https://leetcode.com/problems/decompress-run-length-encoded-list/
package main

func decompressRLElist(nums []int) []int {
	//1-Brute-force approach
	res := make([]int, 0)
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res

}

//func main() {
//	arr := []int{1, 1, 3, 4}
//	nums := decompressRLElist(arr)
//	fmt.Println(nums)
//}
