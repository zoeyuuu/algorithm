package stack

import "fmt"

// 20. 有效的括号 easy
// https://leetcode.cn/problems/valid-parentheses/
// 用栈实现括号的匹配

func Problem20() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	//遇到左括号一直是加入栈不用判断；右括号时需要判断是否出错 故map的键要设置成右括号
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	// for range str是rune字符类型;普通for循环s[i]是byte字节类型
	for i := 0; i < len(s); i++ {
		//键存在（为右括号）
		if pairs[s[i]] > 0 {
			//栈为空 或者 括号不匹配
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			//出栈
			stack = stack[:len(stack)-1]
		} else { //左括号
			stack = append(stack, s[i])
		}
	}
	// 考虑遍历完栈不空的情况
	return len(stack) == 0
}
