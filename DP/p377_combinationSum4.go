package DP

import (
	"fmt"
)

// 377. 组合总和 Ⅳ medium
// https://leetcode.cn/problems/combination-sum-iv/description/
// 完全背包
// 计算装满背包的方式时 遍历顺序有关
// **如果求组合数就{1，2}是外层for循环遍历物品（物品是有序的）**  p518
// **如果求排列数{1，2}{2，1}就是外层for遍历背包（每次循环迭代背包容量都经过每个物品的计算 所以是无序的）** p377 本题

func Problem377() {
	nums := []int{1, 2, 3}
	target := 4
	fmt.Println(combinationSum4(nums, target))
}

// 时间复杂度: O(target * n)，其中 n 为 nums 的长度
// 空间复杂度: O(target)
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	// 求装满背包的排列 先遍历背包容量
	// 避免符号错乱 外层用j内层用i
	for j := 0; j <= target; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				// 装满背包方式 +=
				dp[j] += dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
