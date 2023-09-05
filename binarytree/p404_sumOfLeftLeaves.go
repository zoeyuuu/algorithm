package binarytree

import "fmt"

// 404. 左叶子之和 easy
// https://leetcode.cn/problems/sum-of-left-leaves/
func Problem404() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(sumOfLeftLeaves(root))
}
func sumOfLeftLeaves(node *TreeNode) int {
	if node == nil {
		return 0
	}
	sum := sumOfLeftLeaves(node.Left) + sumOfLeftLeaves(node.Right)
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		return node.Left.Val + sum
	}
	return sum
}
