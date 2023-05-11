//Ques:- https://leetcode.com/problems/construct-string-from-binary-tree/
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	left := ""
	right := ""
	if root.Left == nil && root.Right != nil {
		left = "()"
		right = "(" + tree2str(root.Right) + ")"
	} else if root.Left != nil && root.Right == nil {
		left = "(" + tree2str(root.Left) + ")"
	} else if root.Left != nil && root.Right != nil {
		left = "(" + tree2str(root.Left) + ")"
		right = "(" + tree2str(root.Right) + ")"
	}

	return fmt.Sprintf("%d%s%s", root.Val, left, right)
}
