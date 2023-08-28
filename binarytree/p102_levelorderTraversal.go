package binarytree

import "fmt"

// 层序遍历
// https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
// 可用于解决？：求节点的双亲、孩子节点、二叉树深度、叶子节点个数、判断两棵二叉树是否相等

func Problem102() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = nil
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res := levelOrderTraversal(root)
	fmt.Println(res)
}

// 层序遍历
func levelOrderTraversal(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root} //根节点初始化队列
	// 外层循环用来遍历二叉树的层级
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{}) //创建一个新的空的整数切片 存当前层
		p := []*TreeNode{}         // 建立一个新队列用于下一层
		// 内层for循环用于当前层所有节点
		for j := 0; j < len(q); j++ {
			node := q[j]
			res[i] = append(res[i], node.Val)
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
