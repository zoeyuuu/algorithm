package stack_queue

import (
	"fmt"
	"strconv"
)

// 394. 字符串解码 2023-11-09 67
// https://leetcode.cn/problems/decode-string/description/?envType=study-plan-v2&envId=top-100-liked
// 用栈实现 遇到非']'元素入栈
// 遇到']'用指针从栈顶往下寻找 分别找到'['左右的str,num 出栈到数字前！！！，然后写一个循环将num个str压入栈中
// 最后返回栈
// 注意点1.循环num压[]入栈时 不能用切片 否则会有底层引用问题 先转换为string
// 2.数字可能为多位数 写一个循环k找到
// 3.判断byte是否为数字 可以用stack[k] <= '9' && stack[k] >= '0' 或者 用unicode.IsDigit(r rune) 但得是rune类型 强转或者用for range

func Problem394() {
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
}

// 自己逻辑写的 调试几次ac
// 全程用byte操作 相比其他答案写的不错
func decodeString(s string) string {
	var stack []byte
	//str := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == ']' {
			j := len(stack) - 1
			// j指针找到'['
			for j >= 0 && stack[j] != '[' {
				j--
			}
			// stack[j+1:]转换成string类型 不然切片在重复添加的时候会有底层引用问题
			str := string(stack[j+1:])
			// 写错数字可能是多个 23 所以应该用循环k找 num
			k := j - 1
			// k指针找到数字之前的byte位置
			for k >= 0 && stack[k] <= '9' && stack[k] >= '0' {
				k--
			}
			num, _ := strconv.Atoi(string(stack[k+1 : j]))
			stack = stack[:k+1]
			for num != 0 {
				stack = append(stack, str...)
				num--
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
