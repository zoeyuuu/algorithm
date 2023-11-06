package backtracking

import "fmt"

// 131. 分割回文串 meidum
// https://leetcode.cn/problems/palindrome-partitioning/description/
// 回溯算法 与39/40/77相同

func Problem131() {
	s := "bbbb"
	fmt.Println(partition(s))
}

func partition(s string) (res [][]string) {
	n := len(s)
	if n == 0 {
		return
	}
	// i下标之前为分割点
	path := []string{}
	var dfs func(str string)
	dfs = func(str string) {
		// 终止条件是str==""
		if len(str) == 0 {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		// 当str长度为1时 dfs(str[i:])为“”
		for i := 1; i < len(str)+1; i++ {
			if isPalindrome(str[:i]) {
				path = append(path, str[:i])
				dfs(str[i:])
				path = path[:len(path)-1]
			}
		}
	}
	dfs(s)
	return
}
func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
