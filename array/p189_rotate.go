package array

// 189. 轮转数组 medium
// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
// https://leetcode.cn/problems/rotate-array/description/?envType=study-plan-v2&envId=top-100-liked

func Problem189() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotateArray(nums, 3)
}

// 使用辅助数组
func rotateArray(nums []int, k int) {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[(i+k)%n] = nums[i]
	}
	copy(nums, ans)
}

// 原地反转 反转三次 分界下表是多少模拟一下
func rotateArray2(nums []int, k int) {
	n := len(nums)
	// 写错：k可能大于n先模n一下 否则有些用例过不了
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, n-k-1)
	reverse(nums, n-k, n-1)
}
func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
