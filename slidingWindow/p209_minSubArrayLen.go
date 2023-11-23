package slidingWindow

import (
	"fmt"
	"math"
)

// 209. 和≥target长度最小的子数组 2023-08-10 49
// https://leetcode.cn/problems/minimum-size-subarray-sum/

func Problem209() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen_BF(target, nums))
	fmt.Println(minSubArrayLen_smoothWindow(target, nums))
}

// 滑动窗口法
// 使用两个指针 start 和 end 来表示一个滑动窗口 开始时，两个指针都指向数组的第一个元素。
// 然后，我们通过移动 end 指针，不断扩大窗口，直到窗口内元素的和大于或等于 target 为止。
// 此时，我们记录当前窗口的长度，并尝试将窗口缩小（移动 start 指针），使得窗口内元素的和尽可能小，但仍然保持大于或等于 target。
// 时间复杂度O(n)
func minSubArrayLen_smoothWindow(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	lens := math.MaxInt32
	sum := 0
	for start, end := 0, 0; end < n; end++ {
		sum += nums[end]
		for sum >= target {
			lens = Min(lens, end-start+1)
			sum -= nums[start]
			start++
		}
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
				lens = Min(j-i+1, lens)
				break
			}
		}
	}
	if lens == math.MaxInt32 {
		return 0
	}
	return lens
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
