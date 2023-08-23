package binarytree

// 144 二叉树的中序遍历 easy

import "fmt"

func Problem94() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res1 := inorderTraversal1(root)
	res2 := inorderTraversal2(root)
	fmt.Println(res1, res2)
}

// 递归算法
func inorderTraversal1(root *TreeNode) (res []int) {
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

// 中序遍历 迭代法
// 一直向左（左） 直到为空时出栈访问（中） 转向右子树（右）
func inorderTraversal2(root *TreeNode) (res []int) {
	stack := []*TreeNode{} //初始化栈
	p := root              //遍历指针
	// 栈不空或者p不为空时循环
	for p != nil || len(stack) > 0 {
		for p != nil { //p不为空 入栈 一路向左
			stack = append(stack, p)
			p = p.Left
		}
		// 否则 1.出栈 2.访问 3.转向右子树
		p = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, p.Val)
		p = p.Right
	}
	return
}
