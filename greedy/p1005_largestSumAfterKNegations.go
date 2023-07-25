package greedy

import (
	"fmt"
	"math"
	"sort"
)

// 1005. K 次取反后最大化的数组和 easy

func Problem1005() {
	nums := []int{2, -3, -1, 5, -4}
	k := 2
	fmt.Println(largestSumAfterKNegations(nums, k))
}

// 自己做法
// 先将数组排序 优先反转最小的负数
// 如果负数反转完k仍然还大于零 k为偶数则什么也不做 k为奇数则反转最小的数一次即可
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Ints(nums)
	//i<len(nums)不要忘记 不然全为正数的用例会下标越界
	for i := 0; i < len(nums) && nums[i] < 0 && k > 0; i++ {
		nums[i] = -nums[i]
		k--
	}
	if k%2 == 1 {
		// 其实可以比较nums[i]和nums[i+1]得到最小的数 取反转
		// 但是nums[i+1]可能越界 代码也长 还不如sort一下 不损失时间复杂度
		sort.Ints(nums)
		nums[0] = -nums[0]
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

// 随想录方法
// 使用绝对值排序 也行
func largestSumAfterKNegations2(nums []int, K int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums); i++ {
		if K > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			K--
		}
	}

	if K%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}
