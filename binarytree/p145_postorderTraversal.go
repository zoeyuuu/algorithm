package binarytree

// 144 二叉树的后序遍历 easy

import "fmt"

func Problem145() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	res1 := postorderTraversal1(root)
	res2 := postorderTraversal2(root)
	fmt.Println(res1, res2)
}

// 递归遍历
func postorderTraversal1(root *TreeNode) (res []int) {
	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}
	postorder(root)
	return
}

// 后序遍历 迭代法
// 栈顶元素要想出栈并访问 要么右子树为空 要么刚被访问完（此时左子树早就被访问完）
// 这样就能保证 左右根 的后序遍历顺序
// ***在访问一个节点时 栈中节点刚好是该节点的所有祖先 所以该算法可以应用于：求根节点到某一结点的路径/两个节点的最近公共祖先
func postorderTraversal2(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	p := root         // 遍历指针
	var pre *TreeNode // 辅助指针 指向最近访问过的结点
	for p != nil || len(stack) > 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		p = stack[len(stack)-1] //读取栈顶结点
		// 如果右子树为空或者访问完 1.访问栈顶 2.出栈 3.记录访问完的节点 4.重置遍历指针
		if p.Right == nil || p.Right == pre {
			res = append(res, p.Val)
			stack = stack[:len(stack)-1]
			pre = p // 记录刚刚访问完的结点
			p = nil // 结点访问完之后 重置p指针
		} else {
			p = p.Right
		}
	}
	return
}
