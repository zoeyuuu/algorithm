package DP

import "fmt"

// 1049. 最后一块石头的重量 II  (I heap.1046)
// https://leetcode.cn/problems/last-stone-weight-ii/description/
// 思路：(思路展开 题目意思是把石头分为最相近的两堆)
// 可以将这一堆石头分成两堆（heap1和heap2）保证heap1 - heap2最小即可 (最好情况下heap1 == heap2 = sum/2)
// result = heap1 - heap2 = sum - 2*heap2
// 所以转换为求heap2的最大值 heap2 <= sum/2
// 与p416的思路一致 只需求dp[2/sum]即可得到heap2的最大值
// 0/1 背包 按照p416模板

func Problem1049() {
	stones := []int{2, 7, 4, 1, 8, 1}
	fmt.Println(lastStoneWeightII(stones))
}
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, v := range stones {
		sum += v
	}
	target := sum / 2 //下取整
	dp := make([]int, target+1)
	// 0/1 背包
	for i := 0; i < len(stones); i++ {
		for j := target; j >= stones[i]; j-- {
			dp[j] = max(dp[j], dp[j-stones[i]]+stones[i])
		}
	}
	result := (sum - dp[target]) - dp[target]
	return result
}
