//Ques:- https://leetcode.com/problems/decode-xored-array/
package main

func decode(encoded []int, first int) []int {
	//1st-approach
	res := make([]int, 0)
	res = append(res, first)
	for i := 0; i < len(encoded); i++ {
		val := res[i] ^ encoded[i]
		res = append(res, val)
	}
	return res
	//2nd-approach
	//res := make([]int, len(encoded)+1)
	//res[0] = first
	//for i := 0; i < len(encoded); i++ {
	//	res[i+1] = res[i] ^ encoded[i]
	//}
	//return res
}
