package array

import "fmt"

// 238. 除自身以外数组的乘积 meidum
// https://leetcode.cn/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=top-100-liked
// 一般思路是求出所有乘积和 然后除以当前元素
// 1.题目不给用除法 2.考虑0的情况
// 正逆遍历一边数组 求左右前缀乘积

func Problem238() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}

// 时间复杂度O(N)
// 空间复杂度O(N)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	leftProduct := make([]int, n)
	rightProduct := make([]int, n)
	ans := make([]int, n)
	// 求左侧乘积和 注意循环边界 以及 初始值left[0]=0
	leftProduct[0] = 1
	for i := 1; i < n; i++ {
		leftProduct[i] = nums[i-1] * (leftProduct[i-1])
	}
	// 求右侧乘积和
	rightProduct[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rightProduct[i] = nums[i+1] * rightProduct[i+1]
	}
	for i := 0; i < n; i++ {
		ans[i] = leftProduct[i] * rightProduct[i]
	}
	return ans
}
