package binarytree

import (
	"fmt"
)

// 222. 完全二叉树的节点个数 medium
// https://leetcode.cn/problems/count-complete-tree-nodes/description/
// 主要考察利用完全二叉树的性质 提升算法效率 而不是简单的递归

func Problem222() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	fmt.Println(countNodes(root))
	fmt.Println(countNodes1(root))
}

// 利用完全二叉树的性质（左右子树一棵为完全二叉树 一棵为满二叉树）
// 满二叉树可以利用2^h-1来求节点总数（提高算法效率）
// 因为完全二叉树的子树只能是完全二叉树或满二叉树 故分别一直向左及一直向右比较递归次数是否相等即可判断是不是满二叉树 递归次数即为树高
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := root.Left, root.Right
	i, j := 1, 1 //记录递归次数
	for left != nil {
		left = left.Left
		i++
	}
	for right != nil {
		right = right.Right
		j++
	}
	// 递归次数相同 为满二叉树
	if i == j {
		// Go语言中符号 “ ^ ” 不再用于次方，而是表示“按位异或的运算”
		// 次方使用math.Pow (float);或者2^k 使用左移操作:1<<k
		// return int(math.Pow(2, float64(i)) - 1)
		return 1<<i - 1
	}
	// 完全二叉树 递归调用左右子树
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// 后序遍历递归 对于所有二叉树均适用
// 时间复杂度O（N） 空间复杂度O（logN）
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes1(root.Left) + countNodes1(root.Right) + 1
}
