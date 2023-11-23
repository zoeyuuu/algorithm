package DP

import "fmt"

// 718. 最长重复子数组 meidum  2023-08-10 49
// https://leetcode.cn/problems/maximum-length-of-repeated-subarray/
// 动态规划 一次AC
// dp[i][j] ：以下标i - 1为结尾的A，和以下标j - 1为结尾的B，最长重复子数组长度为dp[i][j]
// dp[i][j] = dp[i-1][j-1] + 1 when nums1[i] == nums2[j]

func Problem718() {
	nums1 := []int{1, 2, 3, 2, 1}
	nums2 := []int{3, 2, 1, 4, 7}
	fmt.Println(findLength(nums1, nums2))
}
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	// 或者dp数组设置为m+1*n+1 省去初始化
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if nums1[i] == nums2[j] {
				// 省去初始化 ==0 的时候赋值为1
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j-1] + 1
				}
			}
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dp[i][j])
		}
	}
	return ans
}
