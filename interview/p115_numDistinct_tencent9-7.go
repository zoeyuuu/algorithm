package interview

import "fmt"

// 115. 不同的子序列 hard
// https://leetcode.cn/problems/distinct-subsequences/
// 用动态规划做

func Problem115() {
	s := "rabbbit"
	subString := "rabbit"
	fmt.Println(numDistinct(s, subString))
}

// 思路：统计字符出现的下标 做成hash表
// 然后按照排列组合（递归）和原则 依据子串顺序 获取递增子串
// ac:50/66
func numDistinct(s, sub string) int {
	mp := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]] = append(mp[s[i]], i)
	}
	ans := 0
	var generate func(int, []int)
	generate = func(i int, result []int) {
		if len(result) == len(sub) {
			ans++
			return
		}
		for j := 0; j < len(mp[sub[i]]); j++ {
			if len(result) == 0 || mp[sub[i]][j] > result[len(result)-1] {
				next := append(result, mp[sub[i]][j])
				generate(i+1, next)
			}
		}
	}
	generate(0, []int{})
	return ans
}
