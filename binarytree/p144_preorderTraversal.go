package binarytree

// 144 二叉树的前序遍历 easy

import "fmt"

func Problem144() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res := preorderTraversal(root)
	fmt.Println(res)
}
func preorderTraversal(root *TreeNode) (res []int) {
	var preorder func(*TreeNode) //声明一个函数类型的变量
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return
}
