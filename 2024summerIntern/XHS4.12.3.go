package main

import (
	"fmt"
)

var (
	n, x int
	a    []int
	memo [][][]int
)

func helper(i, x int, extra bool) int {
	if x == 0 {
		return 0
	}
	if i >= n {
		return n + 1
	}
	if memo[i][x][b2i(extra)] != -1 {
		return memo[i][x][b2i(extra)]
	}

	// 不发
	no := helper(i+1, x, extra)
	// 发一次
	one := n + 1
	if x >= a[i]/2 {
		one = helper(i+1, x-a[i]/2, extra) + 1
	}
	// 多发
	multi := n + 1
	if x >= a[i] && !extra {
		multi = helper(i+1, x-a[i], true) + 1
	}

	return min(min(no, one), multi)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}

func main() {
	fmt.Scan(&n, &x)
	a = make([]int, n)
	memo = make([][][]int, n+1)
	for i := range memo {
		memo[i] = make([][]int, x+1)
		for j := range memo[i] {
			memo[i][j] = make([]int, 2)
			for k := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ans := helper(0, x, false)
	if ans == n+1 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
