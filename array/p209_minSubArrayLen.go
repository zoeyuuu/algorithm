package array

import (
	"fmt"
	"leetcode/divideandconquer"
	"math"
)

func Problem209() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen_BF(target, nums))
	fmt.Println(minSubArrayLen_smoothWindow(target, nums))
}

//209. 和≥target长度最小的子数组

// 滑动窗口法
func minSubArrayLen_smoothWindow(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	lens := math.MaxInt32
	start, end := 0, 0
	sum := nums[0]
	for end < n {
		sum += nums[end]
		for sum >= target {
			lens = divideandconquer.min(lens, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if lens == math.MaxInt32 {
		return 0
	}
	return lens
}

// 暴力算法 超出时间限制 方法一的时间复杂度是 O(n^2)
func minSubArrayLen_BF(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	lens := math.MaxInt32
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= target {
				lens = divideandconquer.min(j-i+1, lens)
				break
			}
		}
	}
	if lens == math.MaxInt32 {
		return 0
	}
	return lens
}
