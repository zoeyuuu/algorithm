package binarytree

import (
	"fmt"
	"math"
)

// 98. 验证二叉搜索树 medium
// https://leetcode.cn/problems/validate-binary-search-tree/description/

func Problem98() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	fmt.Println(isValidBST(root))
	fmt.Println(isValidBST1(root))
}

// 中序遍历（好）
// 中序遍历一定是递增的 只需要比较当前节点值是否大于上一个值即可
// 或者记录中序遍历的结果数组 然后判断是否严格递增
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt32
	var inorder func(node *TreeNode) bool
	inorder = func(node *TreeNode) bool {
		// 终止条件
		if node == nil {
			return true
		}
		// 左
		if !inorder(node.Left) {
			return false
		}
		// 根 1.小于等于上一个访问值则false 2.记录当前访问值
		if node.Val <= pre {
			return false
		}
		pre = node.Val
		// 右
		if !inorder(node.Right) {
			return false
		}
		return true
	}
	return inorder(root)
}

// 递归法：（注意不能简单递归判断子树的左右是否比本结点值大小 还要在上一个递归的上下分为内 所以传参要多一个上下界
// 设计一个递归函数 helper(root, lower, upper) 来递归判断
// 判断子树中所有节点的值是否都在 (l,r) 的范围内（注意是开区间）
// 在递归调用左子树时，我们需要把上界 upper 改为 root.val
// 时间复杂度O(n) 空间复杂度O(n)
func isValidBST1(root *TreeNode) bool {
	var helper func(root *TreeNode, lower, upper int) bool
	helper = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		if root.Val <= lower || root.Val >= upper {
			return false
		}
		return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
	}
	// 初始上下限要设置成math.MinInt64 测试用例有[2147483647]这种极大值
	return helper(root, math.MinInt64, math.MaxInt64)
}
