package backtracking

import "fmt"

// 51. N 皇后  2022-08-14 15
// https://leetcode.cn/problems/n-queens/description/
// 回溯法：按行来尝试摆放皇后 每一行摆放完得到结果 存储
// 剪枝条件：
// 1.不能与之前的皇后一列 用columns存储之前的列
// 2.主对角线 主对角线row-col 是相同的 用diagonals1存储
// 3.副对角线 主对角线row+col 是相同的 用diagonals2存储
// 每次回溯处理四个变量 path,columns,diagonals1,diagonals2

func Problem51() {
	n := 1
	fmt.Println(solveNQueens(n))
}

// 声明全局变量用于保存结果
var res [][]string

func solveNQueens(n int) [][]string {
	// 在线编程平台可能会保持运行环境的状态 重置全局变量
	res = [][]string{}
	path := make([]int, 0, n)
	// 列
	columns := make(map[int]bool)
	// 主对角线()
	diagonals1 := make(map[int]bool)
	// 副对角线()
	diagonals2 := make(map[int]bool)
	dfsNQueens(n, 0, path, columns, diagonals1, diagonals2)
	return res
}

// n是棋盘大小;row是处理的行;path是已经摆放的皇后的列号;columns, diagonals1, diagonals2分别是不能摆放的位置;res是最终结果
func dfsNQueens(n, row int, path []int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == n {
		board := generateBoard(n, path)
		res = append(res, board)
		return
	}
	// i是列号
	for i := 0; i < n; i++ {
		if columns[i] || diagonals1[row-i] || diagonals2[row+i] {
			continue
		}
		path = append(path, i)
		columns[i] = true
		diagonals1[row-i], diagonals2[row+i] = true, true
		dfsNQueens(n, row+1, path, columns, diagonals1, diagonals2)
		delete(diagonals1, row-i)
		delete(diagonals2, row+i)
		delete(columns, i)
		path = path[:len(path)-1]
	}
}

// 生成结果
func generateBoard(n int, path []int) []string {
	var board []string
	// 生成结果的n行
	for i := 0; i < n; i++ {
		// 生成结果的每一行 是一个string类型
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			if j == path[i] {
				row[j] = 'Q'
			} else {
				row[j] = '.'
			}
		}
		board = append(board, string(row))
	}
	return board
}
