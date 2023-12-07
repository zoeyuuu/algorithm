package stack_queue

// 1047. 删除字符串中的所有相邻重复项 easy  2023-08-04 21
// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/description/
// 很简单 用栈实现
func removeDuplicates(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
