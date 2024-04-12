package sort

func sortArray(nums []int) []int {
	// 冒泡排序，比较交换，稳定算法，时间O(n^2), 空间O(1)
	// 每一轮遍历，将该轮最大值放到后面，同时小的往前冒
	// 从而形成后部是有序区
	// compare and swap
	for i := 0; i < len(nums); i++ {
		// 适当剪枝，len()-i到最后的部分都是有序区，避免再排
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}
