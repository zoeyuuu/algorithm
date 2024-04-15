package DP

import (
	"fmt"
	"math"
)

// 122. 买卖股票的最佳时机 I easy
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
// 更好的方法  见greedy 这里写的是动态规划方法

func Problem121() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit1(nums))
}

// 本题dp数组初始化长度就为n （多一个长度省去初始化很难，必须经过一天之后才能持有股票）
// 动态规划 维持两个状态：
// dp[i][0] 表示第i天持有股票所得最多现金;dp[i][1] 表示第i天不持有股票所得最多现金
// 初始化dp[0][0]= dp[0][1] = math.MinInt32 递推式是max 初始化为最小值
// 递推公式：
// 1.dp[i][0] = max(dp[i-1][0], -prices[i]) 由求前一天两个状态得到
// a) dp[i-1][0]前一天就持有股票 第i天继续拥有 现金不变
// b) -prices[i]前一天不拥有股票 第i天购入拥有股票 (本题只能买卖股票一次)
// 2. dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]) 由求前一天两个状态得到
// a) dp[i-1][0]前一天就不持有股票 第i天继续不拥有 现金不变
// b) dp[i-1][0]+prices[i] 前一天拥有股票 第i天卖出股票 (本题只能买卖股票一次)
func maxProfit1(prices []int) int {
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
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[n-1][1]
}

// 贪心 记录已遍历的最小值 然后计算当前值-最小值 求最大值
func maxProfit(prices []int) int {
	result := 0
	min := math.MaxInt32
	for i := range prices {
		min = Min(min, prices[i])
		result = Max(result, prices[i]-min)
	}
	return result
}
