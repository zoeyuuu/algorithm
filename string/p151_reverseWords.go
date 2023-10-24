package string

import "fmt"

// 151. 反转字符串中的单词 medium
// https://leetcode.cn/problems/reverse-words-in-a-string/description/
func Problem151() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s))
}

// 遍历byte切片 用二维切片list存储每一个单词 倒序遍历list添加单词和’ ‘
// 空间复杂度高
func reverseWords(s string) string {
	str := []byte(s)
	var list [][]byte
	var res []byte
	for i := 0; i < len(str); i++ {
		// 遇到空格 跳过 i找到单词的第一个字符
		if str[i] == ' ' {
			continue
		}
		// 不是空格 找到单词
		j := i
		for ; j < len(str) && str[j] != ' '; j++ {
		}
		list = append(list, str[i:j])
		// for循环中已经有i++ i=j即可
		i = j
	}
	for i := len(list) - 1; i >= 0; i-- {
		res = append(res, list[i]...)
		res = append(res, ' ')
	}
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	return string(res)
}

// 方法二：
// 先反转整个字符串（去除冗余的空格） 再翻转每一个单词
// 原地计算 不需要格外的空间
// 挺麻烦的 字符中间的冗余空格很难消除 麻烦 不考虑
func reverseWords2(s string) string {
	return ""
}
