package DP

import "fmt"

// 1143. 最长公共子序列 medium 2023-09-15 120
// 子数组必须连续 子序列不一定要连续
// https://leetcode.cn/problems/longest-common-subsequence/description/
// 一次通过 动态规划
// dp[i][j]：长度为[0, i - 1]的字符串text1与长度为[0, j - 1]的字符串text2的最长公共子序列为dp[i][j]
// 与P72编辑距离有点像 根据text1[i-1] == text2[j-1]来判断
// 如果相等dp[i][j] = dp[i-1][j-1] + 1 如果不等dp[i][j] = max(dp[i-1][j], dp[i][j-1])

func Problem1143() {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	// 省去初始化
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// i,j使用的是个数而非index 要-1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
