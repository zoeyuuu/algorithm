package DP

import (
	"fmt"
	"math"
)

// 322. 零钱兑换 medium
// https://leetcode.cn/problems/coin-change/description/
// 完全背包
// dp下标含义：dp[j]:凑足总额为j所需钱币的最少个数
// 递推式：dp[j] = min(dp[j - coins[i]] + 1, dp[j])
// 初始化：1.凑足总金额为0所需钱币的个数一定是0，那么dp[0] = 0
//	2.递推式是min 防止被覆盖 应该反着 dp[i] = math.MaxInt32 i:1~target
// 遍历顺序： 完全背包：顺序遍历
// 不是组合和排列问题 故外层遍历物品/重量都可

func Problem322() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	// 初始化
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	// 遍历顺序都可 顺序遍历
	for i := 0; i < n; i++ {
		for j := 0; j <= amount; j++ {
			// 递推式
			if j >= coins[i] {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	//没有组合 返回-1
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
