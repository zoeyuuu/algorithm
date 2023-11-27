package sort

import "fmt"

// 归并排序
// 时间复杂度：O(nlogn)  T(n)=2T(n/2)+O(n)
// 空间复杂度：O(n) 空间复杂度即为 1.额外O(n)空间的result数组 2.递归调用的层数最深为logn 最终O(n+logn)=O(n)

func ProblemMergesort() {
	nums := []int{38, 27, 43, 3, 9, 82, 10}
	sortedNums := sortArray(nums)
	fmt.Println(sortedNums)
}

// 对给定数组进行归并排序 返回排序好的数组
// 首先确保左右两个数组都是有序的  然后进行一次归并排序
func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := sortArray(nums[:mid])
	right := sortArray(nums[mid:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	// result用append添加元素 故初始化长度为0
	result := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
