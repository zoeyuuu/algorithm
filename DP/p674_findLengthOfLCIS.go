package DP

// 674. 最长连续递增序列 easy  2023-03-01 12
// https://leetcode.cn/problems/longest-continuous-increasing-subsequence/description/
// 与p300的区别在于要求连续 不需要dp 直接一次遍历即可

func findLengthOfLCIS(nums []int) int {
	maxLen := 1
	curLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curLen++
			maxLen = max(curLen, maxLen)
		} else {
			curLen = 1
		}
	}
	return maxLen
}
