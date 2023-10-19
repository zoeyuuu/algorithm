package DP

import "fmt"

// 70. 爬楼梯 easy
// https://leetcode.cn/problems/climbing-stairs/
// 方法一：斐波那契数列 dp[i] = dp[i-1] + dp[i-2]
// 不考虑dp[0]如何初始化，只初始化dp[1] = 1，dp[2] = 2，然后从i = 3开始递推，这样才符合dp[i]的定义。

func Problem70() {
	n := 45
	fmt.Println(climbStairs(n))
}

// 斐波那契:用一个一维dp数组来保存递推的结果
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

// 改编：一步一个台阶，两个台阶，三个台阶，.......，直到 m个台阶。问有多少种不同的方法可以爬到楼顶呢？
// 1阶，2阶，.... m阶就是物品，楼顶就是背包 可以重复m阶 （完全背包 p377)
// 问跳到楼顶有几种方法其实就是问装满背包有几种方法 且是求排列数->先遍历背包
func climbStairs2(n int) int {
	//定义
	dp := make([]int, n+1)
	//初始化
	dp[0] = 1
	// 本题物品只有两个1,2
	m := 2
	// 遍历顺序
	for j := 1; j <= n; j++ { //先遍历背包
		for i := 1; i <= m; i++ { //再遍历物品 （完全背包 顺序遍历）
			if j >= i {
				dp[j] += dp[j-i]
			}
			//fmt.Println(dp)
		}
	}
	return dp[n]
}
