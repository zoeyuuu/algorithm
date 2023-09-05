package binarytree

import "fmt"

func Problem513() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(findBottomLeftValue(root))
	fmt.Println(findBottomLeftValue1(root))
}

// dfs+前序遍历
// 用curHeight表示遍历的最深层数, curValue记录最左侧值（由于前序遍历先处理左子树，第一次高度增加时一定处理的是左子树）
func findBottomLeftValue(root *TreeNode) int {
	curHeight, curValue := 0, 0
	var dfs func(node *TreeNode, height int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		if height > curHeight {
			curHeight = height
			curValue = node.Val
		}
		dfs(node.Left, height)
		dfs(node.Right, height)
	}
	dfs(root, 0)
	return curValue
}

// 层次遍历 先将右结点加入队列 这样层次遍历是从右往左的
// 遍历的最后一个结点值即为所求
func findBottomLeftValue1(root *TreeNode) int {
	res := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		res = node.Val
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
	}
	return res
}
