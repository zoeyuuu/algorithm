package DP

import "fmt"

// 123 买卖股票的最佳时期3 hard
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/

func Problem123() {
	nums := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(maxProfit3(nums))
}

// 定义五个状态 动态规划
func maxProfit3(prices []int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 5)
	}
	// 五个状态剩余的最大现金 0.未操作1.第一次持有2.第一次不持有3.第二次持有4.第二次不持有
	// 状态0其实可以省略 然后直接dp[i][1] = max(dp[i-1][1], 0-prices[i]) 与122题一样
	dp[0][0], dp[0][1], dp[0][2], dp[0][3], dp[0][4] = 0, -prices[0], 0, -prices[0], 0
	for i := 1; i < n; i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}
	return dp[n-1][4]
}
