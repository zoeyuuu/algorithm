package dfs

import "fmt"

// 463. 岛屿的周长 easy
// https://leetcode.cn/problems/island-perimeter/description/
// 两种方法

func Problem463() {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	//fmt.Println(islandPerimeter_dfs(grid))
	fmt.Println(islandPerimeter_normal(grid))
}

// 方法一 dfs思路：
// 时间复杂度：O(nm)，其中 n 为网格的高度，m 为网格的宽度。每个格子至多会被遍历一次，因此总时间复杂度为 O(nm)。
// 空间复杂度：O(nm)。深度优先搜索复杂度取决于递归的栈空间，而栈空间最坏情况下会达到 O(nm)。
func islandPerimeter_dfs(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 题目只要求一个岛屿 这里考虑多个 通用解法
			if grid[i][j] == 1 {
				perimeter += dfsIsland463(grid, i, j)
			}
		}
	}
	return perimeter
}

// 是周长的情况两种：上一个点是陆地 下一个是1：边界外2：海洋
func dfsIsland463(grid [][]int, i, j int) int {
	inArea := 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
	if !inArea {
		return 1
	}
	if grid[i][j] == 0 {
		return 1
	}
	if grid[i][j] == 2 {
		return 0
	}
	grid[i][j] = 2
	return dfsIsland463(grid, i-1, j) + dfsIsland463(grid, i+1, j) + dfsIsland463(grid, i, j-1) + dfsIsland463(grid, i, j+1)
}

// 方法二 ：常规思路
// 直接遍历 对于每一个=1的点 考虑四个方向 如果是边界外或者是0 那么该条边为周长
// 时间复杂度：O(nm)，每个格子要看其周围 444 个格子是否为岛屿，总时间复杂度为 O(4nm) = O(nm)
// 空间复杂度：O(1) 只需要常数空间存放若干变量
func islandPerimeter_normal(grid [][]int) int {
	perimeter := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 题目只要求一个岛屿 这里考虑多个 通用解法
			if grid[i][j] == 1 {
				perimeter += isPrerimeter(grid, i-1, j) + isPrerimeter(grid, i+1, j) + isPrerimeter(grid, i, j-1) + isPrerimeter(grid, i, j+1)
			}
		}
	}
	return perimeter
}
func isPrerimeter(grid [][]int, i, j int) int {
	if 0 > i || i >= len(grid) || 0 > j || j >= len(grid[0]) || grid[i][j] == 0 {
		return 1
	}
	return 0
}
