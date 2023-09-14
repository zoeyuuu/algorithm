package binarytree

import (
	"fmt"
	"math"
)

// 530 二叉搜索树的最小绝对差 easy
// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/description/
// 与p98模板一样 利用中序遍历 当然也可以存下中序遍历结果再判断

func Problem530() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	fmt.Println(getMinimumDifference(root))
}
func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt32
	pre := math.MinInt32
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if pre != math.MinInt32 {
			ans = min(ans, int(math.Abs(float64(node.Val-pre))))
		}
		pre = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return ans
}
