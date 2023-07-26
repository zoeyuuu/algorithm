package divideandconquer

// 剑指offer第二版 51 hard
// 分治法的归并思想
// 基于归并排序 每次归并排序时需要移动左边数组指针时
// 此时右边数组指针左侧的数与左指针构成逆序数
// 基于归并排序 进行统计
// 时间复杂度O(nlogn) 空间复杂度O(n)

import "fmt"

func Offer51() {
	nums := []int{7, 5, 6, 4}
	k := reversePairs(nums)
	fmt.Println(k)
}

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1) //注意end要-1
}

func mergeSort(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start + (end-start)/2
	cnt := mergeSort(nums, start, mid) + mergeSort(nums, mid+1, end)
	tmp := []int{}
	i, j := start, mid+1
	for i <= mid && j <= end {
		//仅当移动左指针的时候 才有逆序对
		if nums[i] <= nums[j] { //保证相对次序
			tmp = append(tmp, nums[i])
			cnt += j - (mid + 1) // j之前的数都是与nums[i]的逆序数对
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += end - (mid + 1) + 1 //注意这里与之前不同 mid+1~end都是 要+1
	}
	for ; j <= end; j++ {
		tmp = append(tmp, nums[j])
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return cnt
}
