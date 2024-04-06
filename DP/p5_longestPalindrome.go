package DP

import "fmt"

// 5. 最长回文子串 medium
// https://leetcode.cn/problems/longest-palindromic-substring/description/
// 动态规划/中心扩展法 (与p647 回文子串 类似 先看！！！）

func Problem5() {
	s := "babad"
	fmt.Println(longestPalindrome1(s))
	fmt.Println(longestPalindrome2(s))
}

// 方法一：动态规划:dp[i][j]：表示区间范围[i,j]的子串是否是回文子串
// 初始化：dp[i][j] = flase
// dp[i][j]=true 的情况 大前提*** s[i] == s[j]：情况1：j-i<=1 情况2：dp[i+1][j-1] = true
// 遍历顺序：dp[i][j]取决于dp[i+1][j-1] 取决于左下角的值 那么一定是从左下角往右上角遍历 （i: n-1~0 j:i~n-1 j>=i)
// 时间复杂度是 O(N^2) 空间复杂度是 O(N^2)
func longestPalindrome1(s string) string {
	n := len(s)
	maxLen := 0
	ans := ""
	// 初始化值为false
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	// 注意循环起始位置
	for i := n - 1; i >= 0; i-- {
		// j>=i才有意义
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					// 长度判断
					if j-i+1 > maxLen {
						maxLen = j - i
						ans = s[i : j+1]
					}
				} else if dp[i+1][j-1] {
					dp[i][j] = true
					// 长度判断
					if j-i+1 > maxLen {
						maxLen = j - i
						ans = s[i : j+1]
					}
				}
			}
		}
	}
	return ans
}

// 方法二：中心扩展：以中心点往来两边扩展 如果两边字符相等 回文子串++
// 中心2n-1个：单字符中心n , 双字符中心n-1
// 时间复杂度是 O(N^2) 空间复杂度是 O(1)
func longestPalindrome2(s string) string {
	n := len(s)
	// 一个字符的均为回文串
	maxLen := 1
	ans := string(s[0])
	// 单中心扩展
	for i := 0; i < n; i++ {
		left, right := i-1, i+1
		for left >= 0 && right < n {
			if s[left] == s[right] {
				if right-left+1 > maxLen {
					maxLen = right - left + 1
					ans = s[left : right+1]
				}
				left--
				right++
			} else {
				break
			}
		}
	}
	//双中心结点扩展
	for i := 0; i < n-1; i++ {
		//先保证双中心对称
		if s[i] == s[i+1] {
			if 2 > maxLen {
				maxLen = 2
				ans = s[i : i+2]
			}
			left, right := i-1, i+2
			for left >= 0 && right < n {
				if s[left] == s[right] {
					if right-left+1 > maxLen {
						maxLen = right - left + 1
						ans = s[left : right+1]
					}
					left--
					right++
				} else {
					break
				}
			}
		}
	}
	return ans
}
