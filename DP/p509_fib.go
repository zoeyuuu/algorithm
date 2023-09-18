package DP

import "fmt"

// 509. 斐波那契数
// https://leetcode.cn/problems/fibonacci-number/
// dp经典题目

func Problem509() {
	n := 45
	fmt.Println(fib(n))
	fmt.Println(fib1(n))
}

// dp:用一个一维dp数组来保存递推的结果
func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 提高空间效率 不用一维数组记录 只用三个数
func fib1(n int) int {
	if n <= 1 {
		return n
	}
	a := 0
	b := 1
	sum := a + b
	for i := 3; i <= n; i++ {
		a = b
		b = sum
		sum = a + b
	}
	return sum
}
