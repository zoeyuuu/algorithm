package binarytree

import "fmt"

// 226. 翻转二叉树 easy
// https://leetcode.cn/problems/invert-binary-tree/description/
// 用前序或后序遍历递归交换左右子树即可

func Problem226() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = nil
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(invertTree(root))
}
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
