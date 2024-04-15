package DP

import "fmt"

// 96. 不同的二叉搜索树 medium
// https://leetcode.cn/problems/unique-binary-search-trees/description/
// 每一个节点作为头节点的情况的累加
// 这里dp代表的是下标 不是索引 长为n+1
// 状态转移方程： dp[i]= sum(1<=j<=i){dp[j-1]*dp[i-j]}
// 初始化dp[0] = 1(边界节点作为头节点 会有dp[0]的情况 由于递推式是乘法 故为1);dp[1] = 1

func Problem96() {
	fmt.Println(numTrees(3))
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		sum := 0
		// i,j 都表示下标
		for j := 1; j <= i; j++ {
			sum += dp[j-1] * dp[i-j]
		}
		dp[i] = sum
	}
	return dp[n]
}
