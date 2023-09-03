package binarytree

import "fmt"

// 110. 平衡二叉树 easy
// https://leetcode.cn/problems/balanced-binary-tree/
// 利用getHeight高度 递归

func Problem110() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res := isBalanced(root)
	fmt.Println(res)
}
func isBalanced(root *TreeNode) bool {
	h := getHeight(root)
	if h == -1 {
		return false
	}
	return true
}
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := getHeight(root.Left), getHeight(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	return max(left, right) + 1
}
