package interview

/*
字节跳动入营测试 编程题1
题目描述
Excel中，列头标识符是从A开始的字母组成，例如:
A->B->C->...->Z->AA->AB>AC->...->AZ->BA->...BZ--CA->...->ZZ>AAA->....
要求输入一个数字 N,返回从 A到Z 对应的Excel 列头。
注意， 输出的字母要求是大写字母
例如:输入Excel列序号N = 28，输出列头标识符:AB
输入样例: 28 输出: AB
*/

/*
思路：其实excel表示就是26进制的表示，可以将问题抽象成为十进制转换为26进制
每一位表示一个26进制的数字，对应着Excel列头标识符中的一个字母。具体地，'A'对应0，'B'对应1，...，'Z'对应25。
在这个算法中，通过对N不断进行除以26和取余数的操作，来逐步计算每一位的值。将余数对应的字母拼接到结果字符串中，最终得到了Excel列头标识符的表示。
这种将数字转换为一种进制表示的方法，在计算机科学中是常见的。在这个问题中，我们将10进制的数字转换为26进制表示，以便与Excel列头标识符的字母对应起来。
*/
import (
	"fmt"
)

func convertToExcel(n int) string {
	result := ""
	// 进制转换
	for n > 0 {
		//n--和'A'+n%26 为的是对齐 范围是0~25 注意下标
		//或者直接result = string(rune('A'+n%26) -1) + result
		n--
		result = string(rune('A'+n%26)) + result

		// 每次%26计算最低位 /26再计算剩余位
		// 重要**：使用rune()函数将这个整数值转换为对应的Unicode字符
		// 可以正确地用string()表示对应的字母
		n /= 26
	}
	return result
}

func BDcamp1() {
	N := 28
	column := convertToExcel1(N)
	fmt.Println(column)
}

func convertToExcel1(n int) string {
	result := ""
	for n > 0 {
		// +=是右加 所以不能使用
		result = string(rune('A'+n%26-1)) + result
		n /= 26
	}
	return result
}
