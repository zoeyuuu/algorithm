package DP

import "fmt"

func Problem213() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob2(nums))
}

// 213. 打家劫舍 II medium
// https://leetcode.cn/problems/house-robber-ii/description/
// 循环数组 在p198的基础上 分两种情况 1.不包含首元素 2.包含首元素，尾元素不选
func rob2(nums []int) int {
	// 长度为1 单独判断
	if len(nums) == 1 {
		return nums[0]
	}
	return max(rob1(nums[1:]), rob1(nums[:len(nums)-1]))
}
