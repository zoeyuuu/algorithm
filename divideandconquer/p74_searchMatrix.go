package divideandconquer

import (
	"fmt"
	"sort"
)

// 74 搜索二维矩阵 I medium 2023-11-03 36 hot100
// https://leetcode.cn/problems/search-a-2d-matrix/

func Problem74() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	fmt.Println(searchMatrix11(matrix, 20))
	fmt.Println(searchMatrix12(matrix, 20))
}

// 贪心算法
// 等于返回true
// 先往下贪心：如果下一行元素小于target 并且下一行不越界（i < m-1） 则i++
// 不行则往右贪心：如果下一列元素小于target 并且下一列不越界（j < n-1） 则j++
// 其他情况返回flase
func searchMatrix11(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	i, j := 0, 0
	for {
		if matrix[i][j] == target {
			return true
		} else if i < m-1 && matrix[i+1][j] <= target {
			i++
		} else if j < n-1 && matrix[i][j+1] <= target {
			j++
		} else {
			return false
		}
	}
}

// 二分搜索
// 利用两次二分搜索
// 时间复杂度O(logm)+O(logn) = O(logmn)
func searchMatrix12(matrix [][]int, target int) bool {
	// 搜索范围：len(matrix)；返回第一个满足matrix[i][0] > target的下标
	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	// 搜索target的位置 不存在则返回插入的位置
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
