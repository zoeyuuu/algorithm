package main

import (
	"fmt"
	"sort"
)

// 题目：1523 转换成1111 2222 3333 5555的操作数 每次只能+1-1
// 排序+前缀和 可以在O(1)时间内得到操作数
func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	// 复制切片进行操作 最后要根据数组进行输出
	sortedArray := make([]int, n)
	copy(sortedArray, arr)
	pre := make([]int, n+1)
	sort.Ints(sortedArray)
	fmt.Println(sortedArray)
	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + sortedArray[i-1]
	}
	fmt.Println(pre)
	// 排序后利用前缀和可以在O(1)时间内得到操作数
	// cnt := i*a[i]-pre[i] + pre[n]-pre[i+1]-(n-1-i)*a[i]
	mp := make(map[int]int)
	for i := range sortedArray {
		// 去重：首个不跳过+相等跳过
		if i != 0 && sortedArray[i] == sortedArray[i-1] {
			continue
		}
		mp[sortedArray[i]] = i*sortedArray[i] - pre[i] + pre[n] - pre[i+1] - (n-1-i)*sortedArray[i]
	}
	for i := 0; i < n; i++ {
		fmt.Println(mp[arr[i]])
	}
}

// 复杂做法O(n^2)
//func main() {
//	var n int
//	fmt.Scan(&n)
//	arr := make([]int, n)
//	for i := 0; i < n; i++ {
//		fmt.Scan(&arr[i])
//	}
//	mp := make(map[int]int)
//	for i := 0; i < n; i++ {
//		mp[arr[i]] = 0
//	}
//	for value := range mp {
//		sum := 0
//		for i := 0; i < n; i++ {
//			if arr[i] > value {
//				sum += arr[i] - value
//			} else {
//				sum += value - arr[i]
//			}
//		}
//		mp[value] = sum
//	}
//	for i := 0; i < len(arr); i++ {
//		fmt.Println(mp[arr[i]])
//	}
//}
//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}
