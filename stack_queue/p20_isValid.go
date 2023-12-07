package stack_queue

import "fmt"

// 20. 有效的括号 easy 2023-11-30 210
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
	//遇到左括号一直是加入栈不用判断；右括号时要保证栈不空且栈顶为匹配的左括号 否则返回错误
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	// for range str是rune字符类型;普通for循环s[i]是byte字节类型
	for i := 0; i < len(s); i++ {
		// 左括号直接入栈
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			// 栈不空且栈顶为匹配的左括号 出栈
		} else if len(stack) > 0 && stack[len(stack)-1] == pairs[s[i]] {
			stack = stack[:len(stack)-1]
			// 否则返回false
		} else {
			return false
		}
	}
	// 遍历结束栈不空也返回false
	return len(stack) == 0
}
