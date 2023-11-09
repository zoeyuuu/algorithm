package dfs

import "fmt"

// 200. 岛屿数量 medium
// https://leetcode.cn/problems/number-of-islands/description/
// '0'海洋 '1'岛屿 '2'遍历过的岛屿
// 时间复杂度：O(MN)，其中 M 和 N 分别为行数和列数。
// 空间复杂度：O(MN)，在最坏情况下，整个网格均为陆地，深度优先搜索的深度达到 MN

func Problem200() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfsIsland200(grid, i, j)
			}
		}
	}
	return ans
}

// 深度遍历该岛屿的所有地点 改’1‘为’2‘
// dfs终止条件 1：超出边界 2：海洋or遍历过
func dfsIsland200(grid [][]byte, i, j int) {
	inArea := 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
	if !inArea {
		return
	}
	if grid[i][j] != '1' {
		return
	}
	// if grid[i][j] == '1'
	grid[i][j] = '2'
	dfsIsland200(grid, i-1, j)
	dfsIsland200(grid, i+1, j)
	dfsIsland200(grid, i, j-1)
	dfsIsland200(grid, i, j+1)
}
