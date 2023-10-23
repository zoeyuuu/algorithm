package DP

import (
	"fmt"
	"math"
)

// 279. 完全平方数 medium
// https://leetcode.cn/problems/perfect-squares/description/
// 与p322 零钱兑换 完全一样
// 完全背包 dp[j]凑足总额为j所需平方数(钱币)的最少个数
// 递推式：dp[j] = min(dp[j - squares[i]] + 1, dp[j])
// 顺序遍历遍历顺序任意 先物品
// 初始化：dp[0]=0 dp[i]=Math.MaxInt32
// 时间复杂度: O(n * √n)
// 空间复杂度: O(n)

func Problem279() {
	n := 12
	fmt.Println(numSquares(n))
}
func numSquares(n int) int {
	squeres := []int{}
	for i := 0; i*i <= n; i++ {
		squeres = append(squeres, i*i)
	}
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < len(squeres); i++ {
		for j := 0; j <= n; j++ {
			if j >= squeres[i] {
				dp[j] = min(dp[j], dp[j-squeres[i]]+1)
			}
		}
	}
	return dp[n]
}
