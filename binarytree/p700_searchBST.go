package binarytree

import "fmt"

// 700. 二叉搜索树中的搜索 easy
// https://leetcode.cn/problems/search-in-a-binary-search-tree/description/

func Problem700() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	fmt.Println(searchBST1(root, 3).Val)
	fmt.Println(searchBST2(root, 3).Val)
}

// 递归法
func searchBST1(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST1(root.Left, val)
	}
	return searchBST1(root.Right, val)
}

// 迭代法
func searchBST2(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
		} else if root.Val < val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
