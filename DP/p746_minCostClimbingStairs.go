package DP

import "fmt"

// 746. 使用最小花费爬楼梯 easy
// https://leetcode.cn/problems/min-cost-climbing-stairs/description/
// 状态转移方程：dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2])
// dp[i]的LXHY定义：到达第i台阶所花费的最少体力为dp[i]。
// dp[0]= dp[1]= 0 没有实际意义 因为“你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯”

func Problem746() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost)+1)
	n := len(cost)
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
