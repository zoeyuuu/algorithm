package DP

import "fmt"

// 516. 最长回文子序列 medium
// https://leetcode.cn/problems/longest-palindromic-subsequence/description/
//

func Problem516() {
	s := "bbbab"
	fmt.Println(longestPalindromeSubseq(s))
}

// 动态规划
// dp[i][j]表示i~j下标最长回文子序列的长度
// 初始化：dp[i][j] = 0 ; dp[i][i] = 1
// 转移方程dp[i][j] = { 1.s[i] == s[j]时 dp[i][j] = dp[i+1][j-1] + 2 相当于加了两端字符
// 2.s[i] ！= s[j]时 表示同时加入两端字符不行 那考虑分别加入一端字符的最大值dp[i][j] = max(dp[i+1][j], dp[i][j+1]) }
// 遍历顺序：根据状态转移方程 依赖于左、下、左下三个元素 故从左下往右上遍历（i: n-1~0 j:i+1~n-1)

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		// 初始化一个字符为回文子序列的长度
		dp[i][i] = 1
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				// 不用考虑i>j的问题（与p647不同之处）
				// i>j初始化直接为0 直接避免判断
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
