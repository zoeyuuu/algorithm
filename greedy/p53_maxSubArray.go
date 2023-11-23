package greedy

import (
	"fmt"
	"math"
)

// 53 最大子数组和 medium 2023-11-09 264
// https://leetcode.cn/problems/maximum-subarray/description/
// 题目描述：给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。  子数组 是数组中的一个连续部分。

// 解法： 暴力（不能全过）/贪心/动态规划（没写）

func Problem53() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}

// 动态规划
// dp[i]:以下标i元素为结尾的子数组的最大和
// 仔细思考dp[i] = Max(nums[i], dp[i-1]+nums[i]) 就是下面贪心思路的体现
// 注意点1.dp大小+1和i对应是个数还是下标 2.ans := math.MinInt32
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = Max(nums[i], dp[i-1]+nums[i])
	}
	ans := math.MinInt32
	for i := 0; i < n; i++ {
		ans = Max(dp[i], ans)
	}
	return ans
}

// 贪心算法 （说实话有点难想）
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
			result = Max(result, sum)
		}
	}
	return result
}
