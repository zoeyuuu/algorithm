package binarytree

import "fmt"

// 113. 路径总和 II medium
// https://leetcode.cn/problems/path-sum-ii/
// 与p257不同的是 本题递归传递的参数path不是string而是[]int
// 传递切片到函数或将切片分配给另一个变量时，它们实际上是对同一底层数组的引用 切片本质上是一个包装了底层数组的数据结构，它包括一个指向数组的指针、长度和容量
// 故需要手动的对切片进行操作：1.找到结果时 复制切片再append; 2.本层函数运行完需要手动回溯path = path[:len(path)-1]
// path不用作为参数递归传入

func Problem113() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(pathSum(root, 7))
}

// 使用深度遍历 回溯
func pathSum(root *TreeNode, target int) [][]int {
	res := make([][]int, 0)
	path := []int{}
	var travel func(node *TreeNode, targetSum int)
	travel = func(node *TreeNode, targetSum int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		targetSum -= node.Val
		if node.Left == nil && node.Right == nil && targetSum == 0 {
			//  复制切片再append
			copyPath := make([]int, len(path))
			copy(copyPath, path)
			res = append(res, copyPath)
			//复制切片 方法二
			//ans := append([]int(nil), path...)
			//res = append(res, ans)
		}
		if node.Left != nil {
			travel(node.Left, targetSum)
		}
		if node.Right != nil {
			travel(node.Right, targetSum)
		}
		//手动回溯
		path = path[:len(path)-1]
	}
	travel(root, target)
	return res
}
