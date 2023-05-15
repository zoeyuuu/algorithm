package array

import "fmt"

// 59. 螺旋矩阵 II
// 给你一个正整数 n ，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。

func problem59() {
	n := 3
	fmt.Println(generateMatrix(n))
}

type pair struct {
	x, y int
} //结构体表示方向增量

var dirs = []pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} //定义方向：右下左上

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	} //定义及初始化二维切片

	row, col, dirIdx := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[row][col] = i
		dir := dirs[dirIdx] //类型是pair结构体 表示方向增量

		// 坐标超出矩阵边界或者下一个坐标已经被填过
		if r, c := row+dir.x, col+dir.y; r < 0 || r >= n || c < 0 || c >= n || matrix[r][c] > 0 {
			dirIdx = (dirIdx + 1) % 4 //顺时针旋转至下一个方向
			dir = dirs[dirIdx]
		}
		row += dir.x
		col += dir.y
	}
	return matrix
}
