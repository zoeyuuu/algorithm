package backtracking

import (
	"fmt"
	"strconv"
	"strings"
)

// 93. 复原 IP 地址 meidum
// https://leetcode.cn/problems/restore-ip-addresses/
// 两个库函数 广度遍历 3 个即可？

func Problem93() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}
func restoreIpAddresses(s string) []string {
	ans := []string{}
	path := []string{}
	n := len(s)
	var dfs func(index int)
	dfs = func(index int) {
		// // 够四段后就不再继续往下递归 return
		if len(path) == 4 {
			// 此时正好index == n 存放结果
			if index == n {
				str := strings.Join(path, ".")
				ans = append(ans, str)
			}
			return
		}
		// 广度递归
		for i := index; i < n; i++ {
			// 如果有前导0 跳过 05这种形式
			if i != index && s[index] == '0' {
				break
			}
			numStr := s[index : i+1]
			num, _ := strconv.Atoi(numStr)
			// 符合条件深度递归  否则跳过
			if num >= 0 && num <= 255 {
				path = append(path, numStr)
				dfs(i + 1)
				path = path[:len(path)-1]
			} else {
				break
			}
		}
	}
	dfs(0)
	return ans
}
