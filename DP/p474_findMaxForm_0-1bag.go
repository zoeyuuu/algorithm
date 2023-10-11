package DP

import (
	"fmt"
	"strings"
)

// 474. 一和零 medium
// https://leetcode.cn/problems/ones-and-zeroes/
// 0/1背包问题 字符串的个数可以看作是价值 字符串含“0”“1”的个数为重量 重量有两个维度的限制
// dp[i][j][k] 从前i个字符串 取 “0”的数量<=j,”1“的数量<=k 最多的价值（字符串个数）
// i=0 表示不取 i下标:0~len(strs) 共len(strs)+1 i=0 是为了省去初始化
// 状态转移方程：dp[i][j][k] = max(dp[i-1][j][k],dp[i-1][j-w1][k-w2]+1)
//                         = dp[i-1][j][k]  (j<w1 || k<w2)
// 其中w1,w2表示下标为i的字符串含0，1的个数
// 初始化：value ==1 >0 故初始化均为0

// 优化：滚动数组 dp[i][][]的每个元素值的计算只和 dp[i−1][][]的元素值有关 去掉dp的第一个维度
// dp[j][k] = max(dp[j][k], dp[j-w1][k-w2]+1)
// 实现：重量约束的两个内层循环都需要使用**倒序遍历**的形式
// strings.Count(strs[i], "0")统计字符串中子串个数

func Problem474() {
	strs := []string{"10", "0001", "111001", "1", "0"}
	m, n := 5, 3
	fmt.Println(findMaxForm(strs, m, n))
	fmt.Println(findMaxForm2(strs, m, n))
}

// 优化写法 空间复杂度为 O(mn) 时间复杂度：O(lmn+L)
// 二维切片长度m+1,n+1;倒序遍历；
func findMaxForm(strs []string, m int, n int) int {
	l := len(strs)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < l; i++ {
		w1 := strings.Count(strs[i], "0")
		w2 := len(strs[i]) - w1 //strs[i]中“1”的个数
		for j := m; j >= w1; j-- {
			for k := n; k >= w2; k-- {
				dp[j][k] = max(dp[j][k], dp[j-w1][k-w2]+1)
			}
		}
	}
	return dp[m][n]
}

// 标准0/1背包
// 空间复杂度为 O(lmn)  l 是数组 strs的长度 m和n分别是 0和 1的容量
// 时间复杂度：O(lmn+L)  L 是数组 strs中的所有字符串的长度之和
// 三维数组l+1,m+1,n+1;递推关系是i+1 和i 的关系;递推式要判断大小关系
func findMaxForm2(strs []string, m int, n int) int {
	l := len(strs)
	// 三维切片创建 (创建长度为l+1 下表为0的时候为0 表示未选 省去了初始化 这样在循环的时候递推式为 i+1 和i 的关系)
	dp := make([][][]int, l+1) // 0~l
	for i := range dp {
		dp[i] = make([][]int, m+1) //0~m
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1) //0~n
		}
	}
	for i := 0; i < l; i++ {
		w1 := strings.Count(strs[i], "0")
		w2 := len(strs[i]) - w1 //strs[i]中“1”的个数
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				if w1 > j || w2 > k {
					dp[i+1][j][k] = dp[i][j][k]
				} else {
					dp[i+1][j][k] = max(dp[i][j][k], dp[i][j-w1][k-w2]+1)
				}
			}
		}
	}
	return dp[l][m][n]
}
