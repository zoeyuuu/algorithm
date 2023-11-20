package DP

import (
	"fmt"
	"math"
)

// 494. 目标和 medium  2023-09-05 17
// https://leetcode.cn/problems/target-sum/description/
// 方法一：0/1 背包
// 方法二：回溯 （如果要求目标和组合的具体形式 只能用回溯）
// 题意理解：因为是非负整数数组 找到正数组合p + 负数组合q 使得
// p-q=target;p+q=sum 有p=(sum+target)/2
// dp[j](dp[i][j])表示 下标0~i的数字装满大小为j的背包的 组合种类数
// 则转换为0/1 背包 找到大小为p的组合 组合种类数 递推式为dp[j] += dp[j-nums[i]]
// 初始化dp[0] = 1 ( 因为dp[0][0] = 1 )

func Problem494() {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	fmt.Println(findTargetSumWays(nums, target))
}

// 方法一DP
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	p := (sum + target) / 2
	// 排除两类特殊情况 1.所有数加起来都凑不上target 2.(sum+target)%2有余数（凑不齐）
	if sum < int(math.Abs(float64(target))) || (sum+target)%2 == 1 {
		return 0
	}
	dp := make([]int, p+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := p; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[p]
}

// 回溯法 简单 能通过
func findTargetSumWays1(nums []int, target int) (count int) {
	var backtrack func(int, int)
	backtrack = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}
	backtrack(0, 0)
	return
}
