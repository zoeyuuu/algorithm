package DP

import "fmt"

// 746. 使用最小花费爬楼梯 easy
// https://leetcode.cn/problems/min-cost-climbing-stairs/description/
// 状态转移方程：dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])

func Problem746() {

	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	n := len(cost)
	if n == 1 {
		return cost[0]
	}
	if n == 2 {
		return min(cost[0], cost[1])
	}
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
