package binarytree

import "fmt"

// 111 二叉树的最小深度 easy
// https://leetcode.cn/problems/minimum-depth-of-binary-tree/description/
//

func Problem111() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = nil
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(minDepth(root))
	fmt.Println(minDepth2(root))
}

// 深度遍历：递归
// 时间复杂度：O(N) 空间复杂度：O(H) H为树高（递归栈的深度）
func minDepth(root *TreeNode) int {
	// 0.root为空返回0；
	if root == nil {
		return 0
	}
	// 1.左子树为空：返回右子树最小深度+1；2.右子树为空：返回左子树最小深度+1；
	// 3.左右子树都为空（叶子节点）：返回1
	if root.Left == nil || root.Right == nil {
		return minDepth(root.Left) + minDepth(root.Right) + 1 // 简洁写
	}
	// 4.左右子树都不为空：返回子树最小深度 + 1
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// 层序遍历：访问到叶子结点的时候返回层数 即为最小层
// 时间复杂度：O(N)，其中 N 是树的节点数。对每个节点访问一次。
// 空间复杂度：O(N)，其中 N 是树的节点数。空间复杂度主要取决于队列的开销，队列中的元素个数不会超过树的节点数。
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	i := 0
	for ; len(q) > 0; i++ {
		p := []*TreeNode{}
		for j := 0; j < len(q); j++ {
			node := q[j]
			// 叶子节点则返回所在层数
			if node.Left == nil && node.Right == nil {
				return i + 1
			}
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return i + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
