package binarytree

import "fmt"

// 112. 路径总和 easy
// https://leetcode.cn/problems/path-sum/

func Problem112() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(hasPathSum(root, 7))
}

// 深度优先搜索
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 是叶子结点且和为targetSum
	if root.Left == nil && root.Right == nil && targetSum == root.Val {
		return true
	}
	//直接传入targetSum-root.Val完成回溯操作
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
