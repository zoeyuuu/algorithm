package DP

import "fmt"

// 115. 不同的子序列 hard 2023-08-30 17
// https://leetcode.cn/problems/distinct-subsequences/description/
// 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数
// 回溯法过不了 应该用动态规划(写了很久 但写出来了)
// 动态规划要考虑初始化的问题 如果设置dp大小m+1,n+1 那么初始化的时候 dp[i][0] = 1  (numDistinct1)
// 可以直接不考虑省略初始化问题 设置dp大小m,n 然后dp[i][0]和dp[0][i] 分别根据题意初始化  (numDistinct2) 好！

func Problem115() {
	s := "babgbag"
	t := "bag"
	fmt.Println(numDistinct1(s, t))
}

// 动态规划
// dp[i][j]代表s[0~i]中与t[0~j]子序列个数
// 递推式：s[i] == t[j]时 dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
// 否则dp[i][j] = + dp[i-1][j]
// 遍历顺序1~m-1 1~n-1 左上到右下
// 初始化dp[0][i] =0 不用多于操作
// dp[i][0] = sum sum是与t[0]相等的计数
func numDistinct1(s, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	sum := 0
	for i := 0; i < m; i++ {
		if s[i] == t[0] {
			sum++
		}
		dp[i][0] = sum
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// i,j -1 转换为下标
			if s[i] == t[j] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Println(dp)
	return dp[m-1][n-1]
}

// 动态规划
// dp[i][j]代表s[0:i]中与t[0:j]子序列个数
// 递推式须先判断s[i-1] == t[j-1]是否相等
// a) 相等 例如addacd add 此时d相等
// 如果相等 分两种情况相加 --> dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
// 1.使用s[i-1]作为匹配项 个数 = dp[i-1][j-1] 即在addac中匹配ad
// 2.s[i-1]不作为匹配项 个数 = dp[i-1][j] 即在addac中匹配add
// b) 相等 例如addac add 此时c、d不相等 --> dp[i][j] = dp[i-1][j] 即在adda中匹配add
// ***本题最重要的是在初始化 否则就会错***
// dp[0][i]=0 代表在空字符串s中匹配0
// dp[i][0]=1 代表在字符串s中匹配空字符串 为1
func numDistinct2(s, t string) int {
	m, n := len(s), len(t)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		dp[i][0] = 1
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// i,j -1 转换为下标
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Println(dp)
	return dp[m][n]
}

// 回溯法 自己写的
// 超出时间限制 53 / 65 个通过的测试用例
func numDistinct3(s, t string) int {
	path := make([]byte, 0, len(t))
	sum := 0
	var dfs func(index int)
	dfs = func(index int) {
		if len(path) == len(t) {
			sum++
			return
		}
		for i := index; i < len(s); i++ {
			if s[i] != t[len(path)] {
				continue
			}
			path = append(path, s[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return sum
}
