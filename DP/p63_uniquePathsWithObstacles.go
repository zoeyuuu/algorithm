package DP

import "fmt"

// 63. 不同路径 II medium
// https://leetcode.cn/problems/unique-paths-ii/description/
// p62的系列题 区别：1.初始化遇到障碍则第一行/列后序都为0 2.遇到障碍该处dp == 0 不处理

func Problem63() {
	obstacleGrid := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}

// 代码精简版本
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for j := 0; j < n && obstacleGrid[0][j] == 0; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 如果该处是障碍则跳过 跳过本身初值就是0
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 自己写的 0ms
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 初始化 obstacleGrid[i][0] == 1 可以写进循环判断条件里
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 如果该处是障碍则跳过
			if obstacleGrid[i][j] == 1 {
				continue
			}
			// 下面两个判断不需 因为如果是障碍 dp[i][j] == 0 不用考虑
			if obstacleGrid[i-1][j] != 1 {
				dp[i][j] += dp[i-1][j]
			}
			if obstacleGrid[i][j-1] != 1 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
