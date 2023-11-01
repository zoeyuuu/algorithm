package DP

import "fmt"

// 188. 买卖股票的最佳时机 IV hard
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/
// p123题目的一般化 能够买卖k次 就有2k+1个状态 初始化 + 遍历

func Problem188() {
	k := 1
	prices := []int{2, 1}
	fmt.Println(maxProfit4(k, prices))
}
func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2*k+1)
	}
	//初始化
	for i := 1; i < 2*k+1; i += 2 {
		// 奇数 第一天持有股票 初始化为-prices[0]
		dp[0][i] = -prices[0]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < 2*k+1; j++ {
			if j%2 == 0 {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+prices[i])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			}
		}
	}
	return dp[n-1][2*k]
}
