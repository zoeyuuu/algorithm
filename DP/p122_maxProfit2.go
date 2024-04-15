package DP

import (
	"fmt"
)

// 122. 买卖股票的最佳时机 II medium
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
// 更好的方法 贪心 见greedy 这里写的是动态规划方法

func Problem122() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit2(nums))
}

// 动态规划
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] - prices[i])  注意这里是和121. 买卖股票的最佳时机唯一不同的地方。
// dp[i][1] = max(dp[i-1][1], dp[i-1][0] + prices[i]);
// 可以多次买卖 dp[i][0]可以由前一天没持有 购买得到（在这之前可能还买卖过）
func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[n-1][1]
}

// 贪心
func maxProfit(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			sum += (prices[i] - prices[i-1])
		}
	}
	return sum
}
