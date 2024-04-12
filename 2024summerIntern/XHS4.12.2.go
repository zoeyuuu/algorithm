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
	var path []int
	var dfs func(index, sum int, flag bool)
	dfs = func(index, sum int, flag bool) {
		if sum > target {
			return
		}
		if sum == target {
			l := len(path)
			ans = min(ans, l)
			return
		}
		for i := index; i < n; i++ {
			if index > 0 && arr[i] == arr[i-1] {
				continue
			}
			path = append(path, arr[i])
			if flag == true {
				dfs(i+1, sum+arr[i], true)
				dfs(i+1, sum+arr[i]/2, false)
				path = path[:len(path)-1]
			} else {
				dfs(i+1, sum+arr[i], false)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0, 0, true)
	if ans == n+1 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
