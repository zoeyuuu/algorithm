package divideandconquer

import "fmt"

// 35. 搜索插入位置 easy 2023-07-26 6 hot100
// https://leetcode.cn/problems/search-insert-position/description/?envType=study-plan-v2&envId=top-100-liked
// 要求时间的复杂度O(logn) nums为无重复元素的升序排列数组
// 若有重复元素 见p34 searchIndex 返回第一个出现的位置 如果没出现返回插入位置

func Problem35() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}

// 二分搜索 采用左闭右闭区间
// left左边的值一直保持小于target
// right右边的值一直保持大于等于target
// 指针改变时都使用+1 -1
// 最终left = right+1
// 最后返回left
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2 // 防止溢出
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
