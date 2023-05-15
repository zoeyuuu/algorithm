package array

// 33搜索旋转排序数组 medium 要求时间复杂度O(logn)
// 类二分搜索
// 时间复杂度： O(logn)，其中 n 为 nums数组的大小。整个算法时间复杂度即为二分查找的时间复杂度O(logn)。
// 空间复杂度： O(1) 只需要常数级别的空间存放变量。
// 定理一：只有在顺序区间内才可以通过区间两端的数值判断target是否在其中。
// 定理二：判断顺序区间还是乱序区间，只需要对比 left 和 right 是否是顺序对即可，left <= right，顺序区间，否则乱序区间。
// 通过不断的用Mid二分，根据定理二，将整个数组划分成顺序区间和乱序区间，然后利用定理一判断target是否在顺序区间，
// 如果在顺序区间，下次循环就直接取顺序区间，如果不在，那么下次循环就取乱序区间。

import "fmt"

func Problem33() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println()
}
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	left, right := 0, n-1
	for left <= right {
		mid := (left + right) / 2
		switch {
		case nums[mid] == target:
			return mid
		case nums[0] <= nums[mid]: //左侧有序
			if nums[left] <= target && target <= nums[mid] { //target在左侧有序区间 在左侧有序区间查找
				right = mid - 1
			} else { //否则在右侧无序区间查找
				left = mid + 1
			}
		case nums[mid] <= nums[right]: //右侧有序
			if nums[mid] <= target && target <= nums[right] { //target在右侧有序区间
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
