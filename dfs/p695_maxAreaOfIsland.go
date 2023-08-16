package dfs

import (
	"fmt"
)

// 695. 岛屿的最大面积 medium
// https://leetcode.cn/problems/max-area-of-island/description/
// dfs思路 记录每一个岛屿的大小 并选择最大值
// 时间复杂度：O(m×n) 访问每个网格最多一次
// 空间复杂度：O(m×n) 递归的深度最大可能是整个网格的大小

func Problem695() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(grid))
}

func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := dfsIsland695(grid, i, j)
				ans = max(area, ans)
			}
		}
	}
	return ans
}
func dfsIsland695(grid [][]int, i, j int) int {
	inArea := 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
	if !inArea {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}
	// 经常会忘
	grid[i][j] = 2
	// 当前面积1不要忘记加
	return 1 + dfsIsland695(grid, i-1, j) + dfsIsland695(grid, i+1, j) + dfsIsland695(grid, i, j-1) + dfsIsland695(grid, i, j+1)
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
