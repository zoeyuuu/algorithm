/*
小美拿到了一个数组。她定义f(i)为：将第i个元素翻倍后，数组的最大值。现在小美希望你求出f(1)到f(n)的值。你能帮帮她吗？
输入：
5
1 3 2 5 4
输出
5 6 5 10 8
*/
package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int64, n)
	var max int64 = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i] > max {
			max = arr[i]
		}
	}
	for i := 0; i < n; i++ {
		ans := max
		if arr[i]*2 > max {
			ans = arr[i] * 2
		}
		fmt.Print(ans, " ")
	}
}
