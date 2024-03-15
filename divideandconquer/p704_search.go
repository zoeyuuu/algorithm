package divideandconquer

// 二分查找 左闭右闭

func Problem704(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// 写错(left+right)/2
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
