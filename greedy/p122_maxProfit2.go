package greedy

import "fmt"

// 122 买卖股票的最佳时期2 medium
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 方法1.贪心 最适合  2.动态规划 些许复杂 为后续难题服务(见dp maxProfit2)

func Problem122() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit2(nums))
}

// 贪心算法
// 贪心思路：由于不考虑其他因素 利润分解为每天为单位的维度
// 例如第一天买入第四天卖出 可以分解为1-2，2-3，3-4三个小区间
// 局部最优：每一个赚钱间隔都买入+卖出;全局最优：所有赚钱间隔累加
func maxProfit2(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}
