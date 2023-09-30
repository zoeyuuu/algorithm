package DP

import "fmt"

// 647. 回文子串 meidum
// https://leetcode.cn/problems/palindromic-substrings/description/

func Problem647() {
	s := "abc"
	fmt.Println(countSubstrings1(s))
	fmt.Println(countSubstrings2(s))
}

// 方法一：动态规划:dp[i][j]：表示区间范围[i,j]的子串是否是回文子串
// 初始化：dp[i][j] = flase
// dp[i][j]=true 的情况 大前提*** s[i] == s[j]：情况1：j-i<=1 情况2：dp[i+1][j-1] = true
// 遍历顺序：dp[i][j]取决于dp[i+1][j-1] 取决于左下角的值 那么一定是从左下角往右上角遍历 （i: n-1~0 j:i~n-1 j>=i)
func countSubstrings1(s string) int {
	n := len(s)
	result := 0
	// 初始化值为false
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					result++
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					result++
				}
			}
		}
	}
	return result
}

// 方法二：中心扩展：以中心点往来两边扩展 如果两边字符相等 回文子串++
// 中心2n-1个：单字符中心n , 双字符中心n-1
func countSubstrings2(s string) int {
	n := len(s)
	ans := n // 一个字符的均为回文串
	// 单中心扩展
	for i := 0; i <= n; i++ {
		left, right := i-1, i+1
		for left >= 0 && right < n {
			if s[left] == s[right] {
				ans++
				left--
				right++
			} else {
				break
			}
		}
	}
	//双中心结点扩展
}
