package binarytree

import "fmt"

// 199. 二叉树的右视图 medium
// https://leetcode.cn/problems/binary-tree-right-side-view/
// 使用层序遍历 在每一层遍历的时候到q最后一个元素才添加进结果

func Problem199() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = nil
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res := rightSideView(root)
	fmt.Println(res)
}

// 层序遍历
func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		p := []*TreeNode{}
		// 内层for循环用于当前层所有节点
		for j := 0; j < len(q); j++ {
			node := q[j]
			if j == len(q)-1 {
				res = append(res, node.Val)
			}
			// res[i] = append(res[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}
	return res
}
