package string

import "fmt"

// LCR 182. 动态口令 easy
// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/description/
// 把字符串前面的若干个字符转移到字符串的尾部

func Lcr182() {
	password := "s3cur1tyC0d3"
	target := 4
	fmt.Println(dynamicPassword(password, target))
}

// 原地反转 不需要额外的空间复杂度
// 1.反转前半部分2.反转后半部分3.反转整个字符串
func reverseLeftWords(s string, n int) string {
	b := []byte(s)
	// 1. 反转前n个字符
	// 2. 反转第n到end字符
	// 3. 反转整个字符
	reverse(b, 0, n-1)
	reverse(b, n, len(b)-1)
	reverse(b, 0, len(b)-1)
	return string(b)
}

// 切片是引用传递
func reverse(b []byte, left, right int) {
	for left < right {
		b[left], b[right] = b[right], b[left]
		left++
		right--
	}
}

func dynamicPassword(password string, target int) string {
	// 可以直接操作字符串
	return password[target:] + password[:target]
}
