package doublepointer

import "fmt"

// 392. 判断子序列 easy 2022-06-27 6
// https://leetcode.cn/problems/is-subsequence/description/
// 我们初始化两个指针 i 和 j，分别指向 s 和 t 的初始位置,进行匹配
// 最终如果 i 移动到 s 的末尾，就说明 s 是 t 的子序列

func Problem392() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}
