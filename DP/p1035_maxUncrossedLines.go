package DP

import "fmt"

// 1035. 不相交的线 meidum 2023-08-14 3
// https://leetcode.cn/problems/uncrossed-lines/description/
// 与p1143最长公共子序列一模一样！！！
// 判断nums[i]?=nums[i] 相等dp[i][j] = dp[i-1][j-1] + 1
// 不等dp[i][j] = max(dp[i-1][j], dp[i][j-1])

func Problem1035() {
	nums1 := []int{1, 4, 2}
	nums2 := []int{1, 2, 4}
	fmt.Println(findLength(nums1, nums2))
}
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	// 省去初始化
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// i,j使用的是个数而非index 要-1
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
