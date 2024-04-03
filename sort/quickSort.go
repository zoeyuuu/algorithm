package sort

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建一个基于当前时间的随机数生成器
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// 快速排序 Problem912全部重复的时候好像过不了
// 随机化pivot位置提升性能
// 时间复杂度：O(nlogn)
// 空间复杂度：O(logn)

func ProblemQuicksort() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	//rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
func quicksort(nums []int, low, high int) {
	if low < high {
		// 递归对分区点左右两侧进行快速排序
		// 每一次划分确定一个pivot的位置
		pivot := partition(nums, low, high)
		quicksort(nums, low, pivot-1)
		quicksort(nums, pivot+1, high)
	}
}

// 根据pivot进行一次”划分“
func partition(nums []int, low, high int) int {
	// 随机pivot提升性能
	pivotIndex := rng.Intn(high-low+1) + low
	pivot := nums[pivotIndex]
	// i指向的是已经处理完的pivot左侧的最后一个元素(都小于pivot)
	i := low - 1
	for j := low; j <= high-1; j++ { //处理到high前一个位置
		// 遇到比pivot小的元素 放置到i的下一个位置
		if nums[j] < pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 循环完毕之后 i+1是pivot该在的位置
	nums[i+1], nums[high] = nums[high], nums[i+1]
	// 返回pivot下标
	return i + 1
}
