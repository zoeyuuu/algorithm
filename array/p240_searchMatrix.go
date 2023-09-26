package array

import "fmt"

// // 240. 搜索二维矩阵 II medium
// https://leetcode.cn/problems/search-a-2d-matrix-ii/description/
// 1.找到最左上角的元素 2.如果元素值小于target i++ ;大于j-- ;等于 找到 3.超过边界则跳出循环返回false
// 时间复杂度O(M+N)

func Problem240() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(matrix, 20))
}
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
