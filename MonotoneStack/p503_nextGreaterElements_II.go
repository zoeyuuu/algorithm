package MonotoneStack

import "fmt"

// https://leetcode.cn/problems/next-greater-element-ii/description/
// 503. 下一个更大元素 II
// 循环数组 nums 下一个更大元素

func Problem503() {
	nums := []int{1, 2, 3, 4, 3}
	fmt.Println(nextGreaterElements(nums))
	fmt.Println(nextGreaterElements1(nums))
}

// 思路：循环数组 遍历当前元素下一个更大元素最多在当前前一个位置
// 循环遍历两遍数组可以解决 把数组复制一倍模拟可以证明正确性
// 空间复杂度优化 见方法二
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var stack []int
	nums = append(nums, nums...)
	for i, v := range nums {
		for len(stack) != 0 && v > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			// 写错：ans[top%n] = i%n 结果是存的是数字
			ans[top%n] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// 空间复杂度优化
// 避免复制数组 循环2n次 下标用求余方式
func nextGreaterElements1(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	var stack []int
	for i := 0; i < 2*n; i++ {
		v := nums[i%n]
		for len(stack) != 0 && v > nums[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			// 写错：ans[top%n] = i%n 结果是存的是数字
			ans[top] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return ans
}
