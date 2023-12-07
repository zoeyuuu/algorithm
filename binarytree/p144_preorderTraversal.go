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

	res1 := preorderTraversal1(root)
	res2 := preorderTraversal2(root)
	fmt.Println(res1, res2)
}

// 递归
func preorderTraversal1(root *TreeNode) (res []int) {
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

// 前序遍历 迭代法
// 一直
func preorderTraversal2(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		// 一直向左
		for p != nil {
			res = append(res, p.Val) //先访问
			stack = append(stack, p) //入栈
			p = p.Left               //向左
		}
		p = stack[len(stack)-1]      //出栈
		stack = stack[:len(stack)-1] //0~len(stack_queue-2)
		p = p.Right                  //转向右子树 (注意转向右子树的时候都是出栈的结点的右子树)
	}
	return
}
