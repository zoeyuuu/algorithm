package DP

import "fmt"

// 72. 编辑距离 hard
// https://leetcode.cn/problems/edit-distance/description/
// 动态规划dp[i][j]含义：word1(0~i个字符）和word2(0~j)个字符 的编辑距离
// ↑ 表示”个“不表示下标的原因：方便初始化
// 转移方程：  根据word1[i-1] == word2[j-1]是否相等区分 1.如果两个位置字符相等 那编辑距离等于上一个位置dp[i-1][j-1]
// 2.如果不等 那编辑距离肯定得涉及word1[i-1]、word2[j-1] 有三种情况 a.替换操作 b.删除word1[i-1] +dp[i-1][j] c.删除word2[j-1] +dp[i][j-1]
// 初始化：第一行第一列dp[i][0]、dp[0][i] = i (没有字符到i个字符编辑距离为i)

func Problem72() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}
	// i,j表示的是第几个字符
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 状态转移方程
			if word1[i-1] == word2[j-1] { // 下标-1
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}
