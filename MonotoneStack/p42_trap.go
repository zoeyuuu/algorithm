package MonotoneStack

import "fmt"

// hard 42. 接雨水
// https://leetcode.cn/problems/trapping-rain-water/description/

func Problem42() {
	height := []int{4, 2, 0, 3, 2, 5}
	fmt.Println(trap(height))
}

// 方法一：单调栈
// 按照单调栈的处理方式：本来是遍历到大于栈顶元素时 弹出栈顶元素 操作 然后入栈
// 不能简单地找右边第一个更大的元素进行计算 考虑一个特殊情况901234方法失效
// 挺难理解的 看一下链接的图示https://leetcode.cn/problems/trapping-rain-water/solutions/692342/jie-yu-shui-by-leetcode-solution-tuvc/
// 思路： 一层一层接雨水 当遇到遍历元素大于栈顶top的情况 如果top下面还有元素left 那么接雨水区域的宽是left+1~i-1  区域高是h[i]和h[left]的最小值-h[top]
//
//	*遍历v小于等于栈顶元素时 入栈 *遍历v大于栈顶元素时 (看链接的过程图)
//	1.栈内仅有一个元素->出栈 什么也不做
//	2.栈内还有第二个元素left->得到一个可以接雨水的区域，该区域的宽度是 i−left−1，高度是 min(height[left],height[i])−height[top]
func trap(height []int) int {
	stack := []int{}
	sum := 0
	for i, v := range height {
		// v大于栈顶
		for len(stack) != 0 && v > height[stack[len(stack)-1]] {
			// 栈深>=2 接雨水区域
			if len(stack) >= 2 {
				top := stack[len(stack)-1]
				left := stack[len(stack)-2]
				area := (i - left - 1) * (min(height[i], height[left]) - height[top])
				sum += area
				stack = stack[:len(stack)-1]
			} else { //栈深为1
				stack = stack[:len(stack)-1]
			}
		}
		// v <= 栈顶 入栈
		stack = append(stack, i)
	}
	return sum
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
