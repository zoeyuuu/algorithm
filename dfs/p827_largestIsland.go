package dfs

import "fmt"

// 827. 最大人工岛 hard
// https://leetcode.cn/problems/making-a-large-island/description/
// 思路：整个题的思路在p695最大岛屿面积代码上添加
// 第一遍遍历的陆地不在标记为2 而是每一个岛屿一个标记k 并同时统计area 建立k->area的映射mp
// 第二遍遍历,对于每一处海洋 统计与其相连接的陆地计算岛屿面积和
// 时间复杂度、空间复杂度：O(m×n) 同p695
//

func Problem827() {
	grid := [][]int{
		{1, 0},
		{0, 1},
	}
	fmt.Println(largestIsland(grid))
}
func largestIsland(grid [][]int) int {
	//标记方向
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ans := 0
	k := 2
	mp := map[int]int{}
	// 第一遍遍历：获得所有岛屿的area 每个岛屿遍历之后标记为k 且建立k到area的映射mp
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area := dfsIsland827(grid, i, j, k)
				// 防止没有海洋的特殊情况 此时最大面积就是岛屿的最大面积
				ans = max(ans, area)
				mp[k] = area
				k++
			}
		}
	}
	//第二遍遍历：对于每一处海洋 统计与其相连接的陆地计算岛屿面积和
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			//思路：如果是海洋 看起四周的格子 如果1.没越界2.不是海洋3.是没统计过的岛屿 那么newArea+=岛屿面积mp[grid[x][y]]
			if grid[i][j] == 0 {
				newArea := 1
				// 标记已经统计过的岛屿
				exist := map[int]bool{}
				for _, d := range dir4 {
					x, y := i+d.x, j+d.y
					if 0 <= x && x < len(grid) && 0 <= y && y < len(grid) && !exist[grid[x][y]] {
						// map键不存在时 值为0/flase/空，故grid[x][y]=0时 +=0
						// 第二种是exist定义直接初始化 := map[int]bool{0,true} 直接不进入循环
						newArea += mp[grid[x][y]]
						exist[grid[x][y]] = true
					}
				}
				ans = max(ans, newArea)
			}
		}
	}
	return ans
}
func dfsIsland827(grid [][]int, i, j, k int) int {
	inArea := 0 <= i && i < len(grid) && 0 <= j && j < len(grid[0])
	if !inArea {
		return 0
	}
	if grid[i][j] != 1 {
		return 0
	}
	// 岛屿标记
	grid[i][j] = k
	return 1 + dfsIsland827(grid, i-1, j, k) + dfsIsland827(grid, i+1, j, k) + dfsIsland827(grid, i, j-1, k) + dfsIsland827(grid, i, j+1, k)
}
