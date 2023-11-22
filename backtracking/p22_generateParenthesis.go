package backtracking

import "fmt"

// 22. 括号生成 meidum 2023-11-04 107
// https://leetcode.cn/problems/generate-parentheses/description/?envType=study-plan-v2&envId=top-100-liked

func Problem22() {
	n := 3
	fmt.Println(generateParenthesis(n))
}

// 自己写的 一次通过
// index作为递归终止条件
// lNum > 0;rNum > lNum && rNum > 0作为可选左右括号的约束条件
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	path := make([]byte, 0, 2*n)
	var backtrack func(index, lNum, rNum int)
	backtrack = func(index, lNum, rNum int) {
		if index == 2*n {
			temp := make([]byte, len(path))
			copy(temp, path)
			res = append(res, string(path))
			return
		}
		if lNum > 0 {
			path = append(path, '(')
			backtrack(index+1, lNum-1, rNum)
			path = path[:len(path)-1]
		}
		// 右括号剩余数量大于左括号 才能添加右括号
		if rNum > lNum && rNum > 0 {
			path = append(path, ')')
			backtrack(index+1, lNum, rNum-1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, n, n)
	return res
}
