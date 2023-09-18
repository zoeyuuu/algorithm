package DP

import "fmt"

// 70. 爬楼梯 easy
// https://leetcode.cn/problems/climbing-stairs/
// dp[i] = dp[i-1] + dp[i-2]
// 其实就是斐波那契数列
// 不考虑dp[0]如何初始化，只初始化dp[1] = 1，dp[2] = 2，然后从i = 3开始递推，这样才符合dp[i]的定义。

func Problem70() {
	n := 45
	fmt.Println(climbStairs(n))
}

// dp:用一个一维dp数组来保存递推的结果
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
