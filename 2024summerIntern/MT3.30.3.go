package main

import (
	"fmt"
	"strings"
)

/*
小美有两个长度相等的字符串，第一个字符串为 s ，第二个字符串为 t 。
小美每次可以选择一个字符串的一个前缀，然后选择一个字母 c ，将选择的前缀的所有字母都变成 c 。
小美想知道她最少要操作几次可以使得 s 和 t 相等
输入描述：
第一行输入一个长度不超过 10^5 的字符串 s 。
第二行输入一个长度与 s 相等的字符串 t 。
输出描述：
第一行输出一个整数 m 表示答案。
接下来 m 行，每行输出用空格隔开的 i,j,c 表示选择第 i 个字符串的长度为 j 的前缀，将前缀所有字母变成 c 。
输入示例：
aabc
abcc
输出示例：
2
2 3 b
2 2 a
说明：
第1次操作选择第2个字符串的长度为3的前缀，将前缀所有字母变成 'b' ，字符串变成 "bbbc" 。
第2次操作选择第2个字符串的长度为2的前缀，将前缀所有字母变成 'a' ，字符串变成 "aabc" 。
%95!!
*/

func main() {
	var str1, str2 string
	fmt.Scanln(&str1)
	fmt.Scanln(&str2)
	operations := minOperation(str1, str2)
	fmt.Println(len(operations))
	for i := len(operations) - 1; i >= 0; i-- {
		op := operations[i]
		fmt.Printf("%d %d %c\n", op[0], op[1], op[2])
	}
	//for _, op := range operations {
	//	fmt.Printf("%d %d %c\n", op[0], op[1], op[2])
	//}
}
func minOperation(str1, str2 string) [][]interface{} {
	if len(str1) == 1 {
		var operations [][]interface{}
		if str1 != str2 {
			operations = append(operations, []interface{}{1, 1, str2[0]})
			return operations
		}
		return operations
	}
	index := -1
	for i := len(str1) - 1; i >= 0; i-- {
		if str1[i] != str2[i] {
			index = i
			break
		}
	}
	if index == -1 {
		return nil
	}
	operations1 := minOperation(str1[:index], strings.Repeat(string(str1[index]), index))
	operations2 := minOperation(str2[:index], strings.Repeat(string(str2[index]), index))
	if len(operations1) <= len(operations2) {
		operations1 = append(operations1, []interface{}{2, index + 1, str1[index]})
		return operations1
	} else {
		operations2 = append(operations2, []interface{}{1, index + 1, str2[index]})
		return operations2
	}
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
