package stack_queue

import "strconv"

// 150. 逆波兰表达式求值 2023-03-01 10
// https://leetcode.cn/problems/evaluate-reverse-polish-notation/description/
// 求后缀表达式的值 用栈实现 遇到数字入栈 遇到运算符出栈两个数字，计算并入栈

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		// strconv.Atoi字符串转换成整数 err?=nil来判断是数字还是操作符
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}
