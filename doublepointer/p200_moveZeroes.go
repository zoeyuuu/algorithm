package doublepointer

func Problem200() {
	nums := []int{1, 0}
	moveZeroes1(nums)
	moveZeroes2(nums)
}

// 283. 移动零 easy
// https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=top-100-liked
// 简洁 想好每个指针的作用
// right用于遍历 一直++
// left指向已经处理好的数组下一个位置(用于交换) left左侧全是正数
// 如果right遍历到0 继续 遍历到非零 则与left交换
func moveZeroes1(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// 自己写的 没弄清楚两个指针的含义
// 交换元素很麻烦
func moveZeroes2(nums []int) {
	i, j := 0, 0
	for {
		for i < len(nums) && nums[i] != 0 {
			i++
		}
		// j指向非0元素
		for j < len(nums) && nums[j] == 0 {
			j++
		}
		if j >= len(nums) || i >= len(nums) {
			break
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
}
