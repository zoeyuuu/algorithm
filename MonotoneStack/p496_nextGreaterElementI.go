package MonotoneStack

import "fmt"

// https://leetcode.cn/problems/next-greater-element-i/
// 496. 下一个更大元素 I
// 单调栈

func Problem496() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}

// 单调栈
// 理解题意 存储的map 对应的是nums2的元素 nums2中下一个更大的元素
// 在遍历一遍nums1 查找键 存储值
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	mp := make(map[int]int)
	res := make([]int, len(nums1))
	var stack []int
	for i, v := range nums2 {
		for len(stack) != 0 && v > nums2[stack[len(stack)-1]] {
			// 取栈顶下标
			top := stack[len(stack)-1]
			// map 键：当前元素值  值：下一个更大的元素值
			// *写错：mp[nums2[stack[top]]] top已经是栈顶元素(nums2下标)
			mp[nums2[top]] = nums2[i]
			// 出栈
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	for i, v := range nums1 {
		if value, ok := mp[v]; ok {
			res[i] = value
		} else {
			// *写错：没写else 导致覆盖
			res[i] = -1
		}
	}
	return res
}
