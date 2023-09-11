package DP

import "fmt"

// 343. 整数拆分 medium
// https://leetcode.cn/problems/integer-break/
// 动态规划：正整数 i 拆分出的第一个正整数是 j，有两种方案
// 1. 将 i 拆分成 j 和 i−j 的和，且 i−j 不再拆分成多个正整数，此时的乘积是 j×(i−j)
// 2. 将 i 拆分成 j 和 i−j 的和，且 i−j 继续拆分成多个正整数，此时的乘积是 j×dp[i−j]
// 状态转移方程： dp[i]= max(1≤j<i){max(j×(i−j),j×dp[i−j])}
// 初始化dp[0], dp[1] = 1,
// 遍历顺序：从前往后

func Problem343() {
	fmt.Println(numberBreak(10))
}
func numberBreak(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j < i; j++ {
			curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
		}
		dp[i] = curMax
	}
	return dp[n]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
