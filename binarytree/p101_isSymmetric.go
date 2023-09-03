package binarytree

import "fmt"

// 101. 对称二叉树 easy(->medium)
// https://leetcode.cn/problems/symmetric-tree/
// 算法的时间复杂度是 O(n)，因为要遍历 n 个节点
// 空间复杂度是 O(n)，空间复杂度是递归的深度

func Problem101() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = nil
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(isSymmetric(root))
}

// 一开始理解错意思了（理解成了所有子树都要对称） 不是绝对对称
// 所以子函数必须需要两个指针 分别比较左右子树是否对称
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 一开始输入两个root
	return checkSymmetric(root, root)
}

// 以这p,q两个结点为基础 两棵树是否对称
// 如果这两个结点都为空  返回true
// 如果这两个节点都不空 那么值相等并且p.Left, q.Right以及p.Left, q.Right递归满足对称 返回true
func checkSymmetric(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil && q != nil {
		return p.Val == q.Val && checkSymmetric(p.Left, q.Right) && checkSymmetric(p.Left, q.Right)
	}
	return false
}
