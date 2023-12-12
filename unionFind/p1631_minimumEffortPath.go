package unionFind

import (
	"fmt"
	"sort"
)

// 1631. 最小体力消耗路径 medium
// 每日一题 2023.12.11
// 使用并查集 首先初始化m*n的并查集 然后按照边权重(高度差绝对值)从低到高连接结点 如果0和m*n-1联通
// 那么此时添加的边的权重即为所求的最小体力消耗值
// 时间复杂度：O(mnlog(mn))
//				边数：mn 边排序时间O(mnlog(mn))；并查集的时间复杂度为 O(mn⋅α(mn))，其中α为阿克曼函数的反函数。
// 空间复杂度：O(mn)，即为存储所有边以及并查集需要的空间

func Problem1631() {
	heights := [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
	fmt.Println(minimumEffortPath(heights))
}

type edge struct {
	src, des, diff int
}

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	edges := []edge{}
	// 添加边集
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			id := i*n + j
			// 添加竖边
			if i > 0 {
				edges = append(edges, edge{id - n, id, abs(heights[i][j] - heights[i-1][j])})
			}
			// 添加横边
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(heights[i][j] - heights[i][j-1])})
			}
		}
	}
	// 按照边的权重进行排序
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].diff < edges[j].diff
	})
	uf := newUnionFind(m * n)
	for _, v := range edges {
		uf.union(v.src, v.des)
		// 首尾联通即返回当前边权重
		if uf.connected(0, m*n-1) {
			return v.diff
		}
	}
	return 0
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
