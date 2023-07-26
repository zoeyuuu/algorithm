package greedy

import (
	"fmt"
	"math"
)

// 121 买卖股票的最佳时期1 easy
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/

func Problem121() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit12(nums))
}

// 暴力 201 / 212 个通过的测试用例 不行
func maxProfit11(prices []int) int {
	result := 0
	for i := range prices {
		for j := i + 1; j < len(prices); j++ {
			result = Max(result, prices[j]-prices[i])
		}
	}
	return result
}

// 遍历一遍 记录遍历完的最小值 然后更新result
func maxProfit12(prices []int) int {
	result := 0
	min := math.MaxInt32
	for i := range prices {
		min = Min(min, prices[i])
		result = Max(result, prices[i]-min)
	}
	return result
}

func Min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
