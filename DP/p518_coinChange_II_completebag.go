package DP

import "fmt"

// https://leetcode.cn/problems/coin-change-ii/description/
// 518. 零钱兑换 II (I 322)
// 完全背包
// 计算装满背包的方式时 遍历顺序有关
// dp[j] += dp[j-coins[i]]
// **如果求组合数就{1，2}是外层for循环遍历物品（物品是有序的）**  p518 本题
// **如果求排列数{1，2}{2，1}就是外层for遍历背包（每次循环迭代背包容量都经过每个物品的计算 所以是无序的）** p377

func Problem518() {
	amount := 5
	coins := []int{1, 2, 5}
	fmt.Println(change(amount, coins))
}
func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	// 求组合数 先遍历物品
	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++ {
			// 计算装满背包的方式 递推公式 初始化dp[0] = 1
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
