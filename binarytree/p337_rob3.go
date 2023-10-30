package binarytree

import "fmt"

// 337. 打家劫舍 III medium
// https://leetcode.cn/problems/house-robber-iii/description/
// 树形结构的打家劫舍

func Problem337() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = nil
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = nil
	root.Right.Right = &TreeNode{Val: 1}
	fmt.Println(rob3(root))
}
func rob3(root *TreeNode) int {
	value1, value2 := rob1(root)
	return max(value1, value2)
}

// value1选本结点最大值 value2不选本结点最大值
// 返回两个状态（类似dp） 一个是选本结点（该节点以下）的最大值 另一个是不选本结点的最大值
func rob1(node *TreeNode) (value1, value2 int) {
	if node == nil {
		return 0, 0
	}
	a, b := rob1(node.Left)
	c, d := rob1(node.Right)
	// 选本结点 只能加上两个左右子树都不选的情况的最大值
	value1 = node.Val + b + d
	// 不选本结点 取左右子树所有情况的最大值
	value2 = max(a, b) + max(c, d)
	return
}
