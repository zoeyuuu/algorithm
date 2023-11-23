package DP

import (
	"fmt"
)

// 300. 最长递增子序列 medium 2023-09-28 163
// https://leetcode.cn/problems/longest-increasing-subsequence/description/?envType=study-plan-v2&envId=top-100-liked
// 题目要求：给你一个整数数组 nums，找到其中最长严格递增子序列的长度
// 思路一：回溯(暴力) 与p491求所有递增子序列一样 利用回溯 记录最长长度并返回
// 22 / 55 个通过的测试用例 超出时间限制
// 用动态规划做

func Problem300() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

// 动态规划 关键点在于dp[i]的意义 如果单纯指0~i的最长递增序列的长度 那递推公式很难推
// dp[i] 以nums[i]结尾的最长递增序列的长度 dp[i] = max(dp[j] for dp[j]<dp[i) + 1
// 最终遍历一遍dp数组取dp[i]最大值
// 时间复杂度O(n^2)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		Len := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				Len = max(dp[j], Len)
			}
		}
		dp[i] = Len + 1
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

// 回溯 22 / 55
func lengthOfLIS1(nums []int) int {
	n := len(nums)
	maxLen := 0
	var path []int
	var dfs func(index int)
	dfs = func(index int) {
		if len(path) >= maxLen {
			maxLen = len(path)
		}
		for i := index; i < n; i++ {
			// 严格递增要求
			if len(path) > 0 && nums[i] <= path[len(path)-1] {
				continue
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return maxLen
}
