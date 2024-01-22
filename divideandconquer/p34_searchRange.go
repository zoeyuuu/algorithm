package divideandconquer

import (
	"fmt"
	"sort"
)

// 34. 在排序数组中查找元素的第一个和最后一个位置 meidum 2023-12-05 68 hot100
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked
// 要求时间复杂度O(log n)

func Problem34() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	fmt.Println(searchRange1(nums, target))
	fmt.Println(searchRange2(nums, target))
}

// 二分搜索 调用库
// 调用sort.SearchInts(nums, target),sort.Search(len,func)库
func searchRange1(nums []int, target int) []int {
	// 第一次target出现的位置
	idx1 := sort.SearchInts(nums, target)
	if idx1 >= len(nums) || nums[idx1] != target {
		return []int{-1, -1}
	}
	// 第二次target出现的位置
	// 利用的是第一个>target的位置然后-1
	idx2 := sort.Search(len(nums), func(i int) bool { return nums[i] > target }) - 1
	// idx2 := sort.SearchInts(nums, target+1) 一样的效果
	return []int{idx1, idx2}
}

// 二分搜索 调用自己写的函数searchIndex 相当于sort.searchInts库函数
func searchRange2(nums []int, target int) []int {
	idx1 := searchIndex(nums, target)
	if idx1 >= len(nums) || nums[idx1] != target {
		return []int{-1, -1}
	}
	idx2 := searchIndex(nums, target+1) - 1
	return []int{idx1, idx2}
}
func searchIndex(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else { // nums[mid] >= target
			right = mid - 1
		}
	}
	return left
}
