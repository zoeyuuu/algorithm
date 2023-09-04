package binarytree

import (
	"fmt"
	"strconv"
)

// 257. 二叉树的所有路径 easy
// https://leetcode.cn/problems/binary-tree-paths/

func Problem257() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(binaryTreePaths(root))
}

// 自己思路写的
// 前序遍历
// 中：加入结点到path 如果是叶子结点 加入res 左右：递归调用左右孩子节点 path加上"->"
// path+"->"作为参数递归调用的过程中 返回本身实现了回溯
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var preorder func(node *TreeNode, path string)
	preorder = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		// 如果是叶子结点
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
		}
		if node.Left != nil {
			preorder(node.Left, path+"->")
		}
		if node.Right != nil {
			preorder(node.Right, path+"->")
		}
	}
	preorder(root, "")
	return res
}
