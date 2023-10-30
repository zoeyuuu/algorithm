package DP

import "fmt"

// 198. 打家劫舍 medium
// https://leetcode.cn/problems/house-robber/description/

func Problem198() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob1(nums))
}

// dp[i]表示偷到第i家能够偷得的最大金额 （个数而非下标）
// dp[0] = 0 没开始偷 dp[1] = nums[0] 偷到第一家
// 递推公式：dp[i] = max(dp[i - 2] + nums[i], dp[i - 1]);
func rob1(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}
