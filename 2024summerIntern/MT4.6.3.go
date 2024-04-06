package main

/*
小美有一个由 n 个互不相等的正整数构成的数组 a，但她一不小心把 a 弄丢了，他想要重新找到 a。
好在她并不是一无所有，她还记得以下有关 a 的信息：
1. 他完全记得数组 b 的样子，并且 b 是数组 a 删除了某个 a_i 后，剩余的部分做前缀和并打乱的结果。
2. 他完全记得数组 c 的样子，并且 c 是数组 a 删除了某个 a_j 后，剩余的部分做前缀和并打乱的结果。
（保证两次删除的 a_i 和 a_j 不是同一个 a 中的元素）。
请你帮她还原出 a 数组吧。
补充：前缀和指一个数组的某下标之前的所有数组元素的和（包含其自身）。
输入描述：
输入包含三行。
第一行一个正整数 n\ (3 \leq n \leq 10^5)，表示数组 a 的长度。
第二行 n-1 个正整数 b_i\ (1 \leq b_i \leq 10^{14})，表示题中所述数组 b。
第二行 n-1 个正整数 c_i\ (1 \leq c_i \leq 10^{14})，表示题中所述数组 c。
（输入保证有唯一解）
输出描述：输出一行 n 个整数，表示你还原出的 a 数组。
输入示例：
4
8 18 14
15 9 1
输出示例：
1 8 6 4
*/

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr1 := make([]int, n-1)
	arr2 := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&arr1[i])
	}
	sort.Ints(arr1)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&arr2[i])
	}
	sort.Ints(arr2)
	pre := 0
	for i := 0; i < n-1; i++ {
		tmp := arr1[i] - pre
		pre = arr1[i]
		arr1[i] = tmp
	}
	pre = 0
	for i := 0; i < n-1; i++ {
		tmp := arr2[i] - pre
		pre = arr2[i]
		arr2[i] = tmp
	}
	fmt.Println(arr1)
	fmt.Println(arr2)
	ans := make([]int, 0, n)
	for i, j := 0, 0; i < n-1 || j < n-1; {
		if i == n-1 {
			ans = append(ans, arr2[j])
			break
		}
		if j == n-1 {
			ans = append(ans, arr1[i])
			break
		}
		if arr1[i] != arr2[j] {
			if arr1[i] == arr2[j+1] {
				ans = append(ans, arr2[j])
				j++
			}
			if arr2[j] == arr1[i+1] {
				ans = append(ans, arr1[i])
				i++
			}

		}
		ans = append(ans, arr1[i])
		i++
		j++
	}
	for i, num := range ans {
		fmt.Print(num)
		if i < len(ans)-1 {
			fmt.Print(" ")
		}
	}
}
