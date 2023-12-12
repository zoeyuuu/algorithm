package DP

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/maximum-product-subarray/description/?envType=study-plan-v2&envId=top-100-liked
func Problem152() {
	nums := []int{2, -5, -2, -4, 3}
	fmt.Println(maxProduct1(nums))
}

// 由于存在正负 在遍历的同时维护两个量(以i下标为结尾的子数组的最大最小值)iMax,iMin 初始化为1
// 遍历过程中按照当前元素的正负 更新iMax,iMin 一次遍历取iMax的最大值
// 时间复杂度O(N)
func maxProduct1(nums []int) int {
	ans := math.MinInt32
	// 由于是连乘 初始化为1
	iMax, iMin := 1, 1
	for i := 0; i < len(nums); i++ {
		// 如果元素<0 以nums[i]结尾的iMax=iMin*v;iMin=iMax*v
		if nums[i] < 0 {
			tmp := iMax // 暂时保存
			iMax = max(iMin*nums[i], nums[i])
			iMin = min(tmp*nums[i], nums[i])
			// 如果元素>=0 以nums[i]结尾的iMax=iMax*v;iMin=iMin*v
		} else {
			iMax = max(iMax*nums[i], nums[i])
			iMin = min(iMin*nums[i], nums[i])
		}
		ans = max(ans, iMax)
	}
	return ans
}

// 自己的思路 经过多次修改 可以AC 但比较复杂
// 思路先不管正负 计算前缀乘积 遇到0 后续前缀乘积重新计算
// 错误测试用例一：[3,-1,4] 应为4 即前缀和为负数 结果为当前的单个数的情况
//
//	解决：简单的将ans与nums中每个数进行比较
//
// 错误测试用例二：[2,-5,-2,-4,3] 正确应为-2*-4*3=24 而不是2*-5*-2
//
//	解决：错误用例在于两个负数乘积也许逆序处理一遍dp数组 看看能不能ac 成功了
func maxProduct2(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	// dp[0] = 0
	for i := 1; i < n+1; i++ {
		if dp[i-1] == 0 {
			dp[i] = nums[i-1]
		} else {
			dp[i] = dp[i-1] * nums[i-1]
		}
	}
	// 发现错误的测试用例[2,-5,-2,-4,3] 正确应为-2*-4*3=24 而不是2*-5*-2
	// 想到逆序处理一遍dp数组 看看能不能ac 成功了
	dp2 := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		if dp2[i+1] == 0 {
			dp2[i] = nums[i]
		} else {
			dp2[i] = dp2[i+1] * nums[i]
		}
	}
	ans := math.MinInt32
	// 156/190
	for i := 1; i < n+1; i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	for i := 0; i < n; i++ {
		if dp2[i] > ans {
			ans = dp2[i]
		}
	}
	// 错误测试用例[3,-1,4] 应为4 即前缀和为负数 结果为当前的单个数
	// 简单的将ans与nums中每个数进行比较
	// 结果：168/190
	for i := 0; i < n; i++ {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}
