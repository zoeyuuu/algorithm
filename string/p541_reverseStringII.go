package string

import "fmt"

// 541. 反转字符串 II easy
// 给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

func Problem541() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseStr(s, k))
}
func reverseStr(s string, k int) string {
	t := []byte(s)
	n := len(t)
	// 注意这个for循环 很简洁
	for i := 0; i < n; i += 2 * k {
		// i反转起点 j反转终点
		j := i + k - 1
		if j >= n {
			j = n - 1
		}
		reverseIndex(t, i, j)
	}
	return string(t)
}
func reverseIndex(t []byte, i, j int) {
	for a, b := i, j; a < b; {
		t[a], t[b] = t[b], t[a]
		a++
		b--
	}
}
