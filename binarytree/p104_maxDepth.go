package binarytree

import "fmt"

// 104. 二叉树的最大深度 easy
// https://leetcode.cn/problems/maximum-depth-of-binary-tree/
// 平衡二叉树的定义：二叉树的每个节点的左右子树的高度差的绝对值不超过 1，则二叉树是平衡二叉树

func Problem104() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = nil
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(maxDepth1(root))
	fmt.Println(maxDepth2(root))
}

// 深度遍历
// 如果我们知道了左子树和右子树的最大深度 l 和 r,该二叉树的最大深度即为 max(l,r)+1
// 时间复杂度：O(n)
// 空间复杂度：O(height)，其中 height 表示二叉树的高度
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth1(root.Left), maxDepth1(root.Right)) + 1
}

// 广度遍历：层序遍历框架
// 返回最后遍历的那一层即可
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	// for循环外定义i记录层数
	i := 0
	for ; len(q) > 0; i++ {
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	// 最后for循环i会++ 所以最后返回的不用i+1
	return i
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
