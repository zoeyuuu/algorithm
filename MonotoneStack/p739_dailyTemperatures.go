package MonotoneStack

import (
	"fmt"
	"math"
)

// 739. 每日温度 meidum
// https://leetcode.cn/problems/daily-temperatures/description/

func Problem739() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))
}

// 方法一：哈希
// 考虑每一个位置：要找到其位置之后比他大的温度的最小下标
// 暴力：遍历之后的元素 但可能时间为O(N)
// 优化：建立i位置之后的hash表（next数组） 温度---第一次出现的下标 查询大于temperature[i]之后的下标的最小值
// 注意：因为当前位置hash表只与之后的元素有关 故不能一次遍历建立hash表 而应该倒序遍历 实时更新hash表
// next数组的初始化 表示的是temperature[i]之后的下标的最小值 故未遍历之前初始化为Math.MaxInt32
// res[i] = MIn(j-i) while(temperature[i]<temperature[j]&& i<j)
func dailyTemperatures(temperatures []int) []int {
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

// 方法二：单调栈
