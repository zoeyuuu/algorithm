package backtracking

import (
	"fmt"
	"sort"
)

func Problem473() {
	matchsticks := []int{1, 1, 2, 2, 2}
	fmt.Println(makesquare(matchsticks))
}

func makesquare(matchsticks []int) bool {
	n := len(matchsticks)
	// 先逆序排序 减少时间复杂度
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks)))
	// sort.Ints(matchsticks)
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}
	tarEdge := sum / 4
	edges := make([]int, 4)
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == n {
			return true
		}
		for i := 0; i < 4; i++ {
			if edges[i]+matchsticks[index] <= tarEdge {
				edges[i] += matchsticks[index]
				if dfs(index + 1) {
					return true
				}
				edges[i] -= matchsticks[index]
			}
		}
		return false
	}
	return dfs(0)
}
