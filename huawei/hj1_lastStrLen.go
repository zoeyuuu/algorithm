package huawei

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// fmt.Scan() 函数默认情况下会在读取到第一个空白字符（包括空格、制表符和换行符）时停止读取。
// 因此，如果需要读取一整行输入，应该使用 bufio 包中的 Scanner 类型来逐行读取输入
// bufio.Scanner 读取输入时会包含输入行的结尾换行符（\n），因此需要在处理输入时注意去除结尾的换行符

func HJ1() {
	scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan读取一行 返回一个bool值 判断读取输入时发生了错误（如读取到了文件末尾或发生了 I/O 错误）
	if scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		fmt.Println(len(arr[len(arr)-1]))
	}
}
