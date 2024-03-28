package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*数组压缩 相邻的数组合并
输入：[1(1),2(2),3(31),3(42),2(12)]
输出：[1(1),2(2),3(73),2(12)]
重点在于字符串的处理
*/

func main() {
	var input string
	var output string
	fmt.Scan(&input)
	input = input[1 : len(input)-1]
	// 按照多种符号切割
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return r == '(' || r == ')' || r == ','
	})
	// fmt.Println(parts)
	output += "["
	preCount := 0
	// 后一个不同才存结果 否则累计preCount
	for i := 0; i < len(parts); i += 2 {
		if i == len(parts)-2 || parts[i] != parts[i+2] {
			// 存
			num, _ := strconv.Atoi(parts[i])
			count, _ := strconv.Atoi(parts[i+1])
			output += fmt.Sprintf("%d(%d),", num, count+preCount)
			preCount = 0
		} else {
			count, _ := strconv.Atoi(parts[i+1])
			preCount += count
		}
	}
	output = output[:len(output)-1]
	output += "]"
	fmt.Println(output)
}
