package binarytree

import (
	"fmt"
	"math"
)

// 501. 二叉搜索树中的众数 easy
// https://leetcode.cn/problems/find-mode-in-binary-search-tree/description/

func Problem501() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 2}
	fmt.Println(findMode1(root))
	fmt.Println(findMode2(root))
}

// *** （仅适用于Bst）
// 不使用额外空间存储结点值 时间复杂度O(N) 空间复杂度O(1)
// 思路与P530相似 利用二叉搜索树中序遍历的结果必然递增 在遍历过程中比较当前结点值和前一个值
func findMode1(root *TreeNode) []int {
	pre := math.MinInt32
	res := []int{}
	var inorder func(node *TreeNode)
	count, maxCount := 1, 1
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if node.Val == pre {
			// 与前一个值相等频率++
			count++
		} else {
			// 重置频率
			count = 1
		}
		if count >= maxCount {
			// 当前count > maxCount, res重置更新
			if count > maxCount {
				res = []int{node.Val}
			} else {
				// 当前count = maxCount, res加上频率相同的值
				res = append(res, node.Val)
			}
			// 最大频率更新
			maxCount = count
		}
		pre = node.Val
		inorder(node.Right)
	}
	inorder(root)
	return res
}

// （对所有二叉树找众数都适用）中序遍历二叉树 存储数字到频率的map 最后找出最大频率的众数
func findMode2(root *TreeNode) []int {
	mp := make(map[int]int)
	res := []int{}
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		mp[node.Val]++ // 统计元素频率
		travel(node.Left)
		travel(node.Right)
	}
	travel(root)
	maxFrequency := 0
	for _, value := range mp {
		maxFrequency = max(maxFrequency, value)
	}
	for key, value := range mp {
		if value == maxFrequency {
			res = append(res, key)
		}
	}
	return res
}
