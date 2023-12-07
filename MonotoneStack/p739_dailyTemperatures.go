package MonotoneStack

import (
	"fmt"
	"math"
)

// 739. 每日温度 meidum
// https://leetcode.cn/problems/daily-temperatures/description/
// 下一个更高温度出现在几天后

func Problem739() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures1(temperatures))
	fmt.Println(dailyTemperatures2(temperatures))
}

// 方法一：单调栈 单调栈存的是下标
// 单调栈 递增(从栈顶到栈底) （为了进栈满足递增顺序 当遍历元素大于/或等于栈顶时，栈顶元素弹出）
// 遍历元素小于等于栈顶元素时 直接入栈
// 遍历元素大于栈顶元素时 栈顶元素出栈并计算 继续比较下一个栈顶
// 遍历完之后单调栈内剩下的元素是没有找到右边更大的元素 该题要求此类情况为0 即什么也不做 因为切片初始化即为0
func dailyTemperatures1(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var stack []int
	for i, v := range temperatures {
		// 栈空 or 遍历元素小于等于栈顶元素 直接入栈
		// 注意判断条件 temperatures[stack_queue[len(stack_queue)-1]] 第一次写错 stack_queue[len(stack_queue)-1]存的是下标 比较的是元素
		for len(stack) != 0 && v > temperatures[stack[len(stack)-1]] {
			// 取栈顶下标
			top := stack[len(stack)-1]
			// 计算处理
			ans[top] = i - top
			// 出栈
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// 方法二：哈希
// 考虑每一个位置：要找到其位置之后比他大的温度的最小下标
// 暴力：遍历之后的元素 但可能时间为O(N)
// 优化：建立i位置之后的hash表（next数组） 温度---第一次出现的下标 查询大于temperature[i]之后的下标的最小值
// 注意：因为当前位置hash表只与之后的元素有关 故不能一次遍历建立hash表 而应该倒序遍历 实时更新hash表
// next数组的初始化 表示的是temperature[i]之后的下标的最小值 故未遍历之前初始化为Math.MaxInt32
// res[i] = MIn(j-i) while(temperature[i]<temperature[j]&& i<j)
func dailyTemperatures2(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	next := make([]int, 101) //存储0~100
	for i := 0; i < 101; i++ {
		next[i] = math.MaxInt32
	}
	for i := n - 1; i >= 0; i-- {
		next[temperatures[i]] = i
		minIndex := math.MaxInt32
		for j := temperatures[i] + 1; j < 101; j++ {
			if next[j] < minIndex {
				minIndex = next[j]
			}
		}
		if minIndex == math.MaxInt32 {
			res[i] = 0
		} else {
			res[i] = minIndex - i
		}
	}
	return res
}
