package greedy

import (
	"fmt"
	"math"
)

// 53 最大子数组和 medium
// https://leetcode.cn/problems/maximum-subarray/description/
// 题目描述：给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。  子数组 是数组中的一个连续部分。

// 解法： 暴力（不能全过）/贪心/动态规划（没写）

func Problem53() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray1(nums))
}

// 贪心算法
// 原理：当连续和为负数时 从下一个位置重新开始统计 因为负数加上后面只会使连续和变小
// 贪心贪在前序和是正数 那就继续统计 用一个变量result记录当前统计到的最大值 最终返回
func maxSubArray1(nums []int) int {
	result := math.MinInt32
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > result {
			result = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return result
}

// 暴力解法 204 / 210 个通过的测试用例 其余超出时间限制
func maxSubArray2(nums []int) int {
	result := math.MinInt32 //要设置成最小值不能设置成0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = 0 //内层循环完sum需要清零
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			result = max(result, sum)
		}
	}
	return result
}
