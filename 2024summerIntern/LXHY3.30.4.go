package main

import (
	"fmt"
	"sort"
)

/*
对于 30% 的数据: N ≤ 100; M ≤ 100.
对于 50% 的数据: N ≤ 1000; M ≤ 1000.
对于 70% 的数据: N ≤ 10000; M ≤ 10000.
对于 100% 的数据: 0 < N ≤ 10^6; 0 < M ≤ 10^5; 0 ≤ Fi < 2^64; 0 < l < r ≤ N.
*/
func main() {
	var m, n int
	fmt.Scan(&m, &n)
	arr := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&arr[i])
	}
	var left, right int
	var ans string
	for i := 0; i < n; i++ {
		fmt.Scan(&left, &right)
		ans += isTriangle(arr[left-1 : right])
	}
	fmt.Println(ans)
}
func isTriangle(arr []int) string {
	if len(arr) < 3 {
		return "N"
	}
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	for i := 0; i < len(sorted)-2; i++ {
		for j := i + 1; j < len(sorted)-1; j++ {
			if sorted[i]+sorted[j] > sorted[j+1] {
				return "Y"
			}
		}
	}
	return "N"
}
