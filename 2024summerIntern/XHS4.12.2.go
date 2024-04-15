package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, target int
	fmt.Scan(&n, &target)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	ans := n + 1
	var dfs func(index, sum, num int, flag bool)
	dfs = func(index, sum, num int, flag bool) {
		if sum > target {
			return
		}
		if sum == target {
			ans = min(ans, num)
			return
		}
		for i := index; i < n; i++ {
			if i > index && arr[i] == arr[i-1] {
				continue
			}
			if flag == true {
				dfs(i+1, sum+arr[i], num+1, true)
				dfs(i+1, sum+arr[i]/2, num+1, false)
			} else {
				dfs(i+1, sum+arr[i], num+1, false)
			}
		}
	}
	dfs(0, 0, 0, true)
	if ans == n+1 {
		ans = -1
	}
	fmt.Println(ans)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
