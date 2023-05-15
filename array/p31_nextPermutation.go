package array

// 31. 下一个排列 medium
// 整数数组的 下一个排列 是指其整数的下一个字典序更大的排列
// 思路：其实就是找组合中比当前数大的下一个数
// 原理：从数组尾部向前扫描 得到第一个 顺序的 nums[j] < nums[j+1] 此时nums[j]就是需要交换的位置
//      为了让交换后数组比原来大一点点 需要在j+1之后的逆序数组中找到比nums[j]大一点点的位置k 交换nums[j]和nums[k]
//      并且将交换之后 j+1之后的逆序数组转置成顺序 以保证值最小
// 时间复杂度O(N) 遍历加上reverse都是O(N)
// 空间复杂度O(1)

import (
	"fmt"
)

func Problem31() {
	nums := []int{1, 1, 5, 1, 1}
	nextPermutation(nums)
	fmt.Println(nums)
}

func nextPermutation(nums []int) {
	n := len(nums)
	// 长度小于等于一不用交换元素排序等 单独考虑
	if n <= 1 {
		return
	}
	//n=2 看看能不能合并 以防万一写一下比较好
	if n == 2 {
		nums[0], nums[1] = nums[1], nums[0]
		return
	}
	// 找到 得到第一个 顺序的 nums[j] < nums[j+1]下标j
	// j = -1 则代表全部逆序 不需要交换 直接reverse nums[j+1:]
	j := n - 2
	for ; j >= 0 && nums[j] >= nums[j+1]; j-- { // >=；注意边界条件j会到-1
	}

	// j >= 0 交换
	if j >= 0 {
		k := n - 1 //找到最小的nums[k]>nums[j]
		// <=；注意从尾部往前k--不需要考虑越界问题
		for nums[k] <= nums[j] && k > j {
			k--
		}
		nums[j], nums[k] = nums[k], nums[j] //交换nums[j]和nums[k]
	}

	// sort.Ints(nums[j+1:]) 可以调用但本身顺序可以节省时间复杂度
	// reverse
	for i, j := j+1, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
