/*
小美拿到了一个数组。她定义f(i)为：将第i个元素翻倍后，数组的最大值。现在小美希望你求出f(1)到f(n)的值。你能帮帮她吗？
输入：
5
1 3 2 5 4
输出
5 6 5 10 8
输入描述：
第一行输入一个正整数n，代表数组大小。
第二行输入n个正整数a_i，代表小红拿到的数组。
1\leq n \leq 200000
1\leq a_i \leq 10^9

30%
 */

package main

import (
	"fmt"
	"strings"
)

package main

import (
"fmt"
"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	var max int = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		if arr[i] > max {
			max = arr[i]
		}
	}
	var result strings.Builder
	for i := 0; i < n; i++ {
		if arr[i]*2 >max{
			result.WriteString(fmt.Sprintf("%d ",arr[i]*2))
		}
		result.WriteString(fmt.Sprintf("%d ",max))
	}
	fmt.Println(result.String())
}
