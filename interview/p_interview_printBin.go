package interview

import (
	"fmt"
	"strings"
)

// 2023.3.2
// 面试题 05.02. 二进制数转字符串

// 十进制小数转二进制小数的方法是：小数部分乘以 2，取整数部分作为二进制小数的下一位，小数部分作为下一次乘法的被乘数，
// 直到小数部分为 0 或者二进制小数的长度超过 32 位。

func problem_printBin() {
	fmt.Println(printBin(0.5))
}
func printBin(num float64) string {
	sb := &strings.Builder{} // 创建一个 strings.Builder 类型的变量
	sb.WriteString("0.")     // 写入 "0."，表示二进制数的小数点前的部分
	for sb.Len() <= 32 && num != 0 {
		num *= 2
		digit := byte(num) // 结果转换为byte类型的整数（0或1）

		if digit != 0 {
			sb.WriteString("1")
		} else {
			sb.WriteString("0")
		}

		//或者更简洁的写法 该数字的ASCII码值加上'0'（48）的结果写入字符串生成器sb中
		//sb.WriteByte('0' + digit)

		num -= float64(digit)
	}
	if sb.Len() <= 32 {
		return sb.String()
	}
	return "ERROR"
}
