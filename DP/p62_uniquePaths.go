package DP

import (
	"fmt"
	"math/big"
)

// 62. 不同路径 medium
// https://leetcode.cn/problems/unique-paths/

func Problem62() {
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths1(3, 7))
	fmt.Println(uniquePaths2(3, 7))
}

// 动态规划 dp[i][j] = dp[i-1][j] + dp[i][j-1]
// dp[i][0] = dp[0][j] = 1
func uniquePaths(m, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := range dp[0] {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 组合数 相当于总共移动m-1次向下，n-1次向右
// ans = ${C_{m-n-2}}^{n-1}$ 总共有这么多种可能
// 1.利用库函数 2.
// {C_{m-n-2}}^{n-1} = (m-n-2)!/(n-1)!*(m-1)! = (m-n-2)*...*m/(n-1)*...1
func uniquePaths1(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
func uniquePaths2(m, n int) int {
	ans := 1
	for x, y := n, 1; y < m; x, y = x+1, y+1 {
		ans = ans * x / y
	}
	return int(ans)
}
