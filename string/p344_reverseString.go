package string

import "fmt"

// 344. 反转字符串 easy
// https://leetcode.cn/problems/reverse-string/description/

func Problem344() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	// string()将字节切片转换为字符串
	fmt.Println(string(s))
}
func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
