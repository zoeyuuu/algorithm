package dfs

import "fmt"

// 79. 单词搜索 meidum 2023-10-24 38
// https://leetcode.cn/problems/word-search/description/?envType=study-plan-v2&envId=top-100-liked
// 用index判断是否匹配完毕
// 约束条件1.值等于word[index] 2.在board范围内
// 对于标记访问过的位置 可以将此处值改为'/'然后深度遍历之后 回溯改回来(就不用使用一个visited数组了)

func Problem79() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCCED"
	fmt.Println(exist(board, word))
}
func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	// direction代表上一层的方向
	// flag标记是否找到
	flag := false
	var backtrack func(index, row, col int)
	backtrack = func(index, row, col int) {
		// flag == true直接返回 不用继续找了
		if flag || index == len(word) {
			flag = true
			return
		}
		// 越界判断
		if row < 0 || row >= m || col < 0 || col >= n {
			return
		}
		if board[row][col] == word[index] {
			temp := board[row][col]
			// 标记访问过的位置
			board[row][col] = '/'
			backtrack(index+1, row, col+1) // 向右
			backtrack(index+1, row+1, col) // 向下
			backtrack(index+1, row, col-1) // 向左
			backtrack(index+1, row-1, col) // 向上
			board[row][col] = temp         //回溯还原
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			backtrack(0, i, j)
		}
	}
	return flag
}
