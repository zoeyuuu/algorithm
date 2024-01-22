package divideandconquer

// 162 寻找峰值 2023-12-09 62
// https://leetcode.cn/problems/find-peak-element/
// 必须实现时间复杂度为 O(logn)的算法
// 相邻元素不相等

func Problem162() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	findPeakElement1(nums)
	findPeakElement2(nums)
}

// 二分查找 比较中间mid、mid+1选择较大的元素一边缩小搜索范围
// 时间复杂度 O(logn)
func findPeakElement1(nums []int) (idx int) {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			right = mid
		}
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		}
	}
	return left
}

// 遍历一遍 寻找最大的元素
// 时间复杂度O(n) 不满足题意 但能AC
func findPeakElement2(nums []int) (idx int) {
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return
}
