package graph

import "fmt"

// 每日一题 1466. 重新规划路线
// https://leetcode.cn/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/description/
// 其实给定的是一颗“树” 其中有一些逆序边 题目要求统计需要转置的边的数量 使得所有节点都能到达根节点0

func Problem1466() {
	n := 6
	connections := [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}}
	fmt.Println(minReorder1(n, connections))
}

// 用邻接表存储图信息 节省时间复杂度
// 不能用单纯的有向图邻接表存储方式 因为这样遇到正向的边就遍历不到
// （有向图的邻接表遍历从一个结点开始只能遍历一个连通分量 而本题即使不连通 也要遍历到(逆向)
// 思路：采用无向图的邻接表存储方式 同一条边存储两遍 但是存储变得同时还需要存储边的方向
// 0代表由其他节点到“0”的方向 1代表需要改边的边
// 即存储每一个节点的所有邻居 然后深度遍历 边的去重使用dfs函数中的父节点这一参数来实现
// 时间复杂度 O(N) 空间复杂度O(N)
func minReorder1(n int, connections [][]int) int {
	ans := 0
	adj := make([][][]int, n)
	// 构建“邻接表”
	for _, edge := range connections {
		adj[edge[0]] = append(adj[edge[0]], []int{edge[1], 1})
		adj[edge[1]] = append(adj[edge[1]], []int{edge[0], 0})
	}
	var dfs func(parent, x int)
	dfs = func(parent, x int) {
		// 这里edge[0]是终点边 edge[1]的值代表方向
		for _, des := range adj[x] {
			if des[0] == parent {
				continue
			}
			ans += des[1]
			dfs(x, des[0])
		}
	}
	// e1:父节点 e2:遍历节点
	dfs(-1, 0)
	return ans
}

// 自己写的方法 71/76
// 用一个visited数组标记访问过的边 从0开始深度遍历 遇到含有0的边继续
// 0为起点 表示这条边需要reverse ans++ 递归dfs(边终点)
// 0为终点 表示这条边无需操作 递归dfs(边终点)
// 但是由于递归每一次都需要遍历整个边 所以时间复杂度超出
func minReorder2(n int, connections [][]int) int {
	visited := make([]bool, len(connections))
	ans := 0
	var dfs func(num int)
	dfs = func(num int) {
		for i, v := range connections {
			if v[0] == num && !visited[i] {
				visited[i] = true
				ans++
				dfs(v[1])
			}
			if v[1] == num && !visited[i] {
				visited[i] = true
				dfs(v[0])
			}
		}
	}
	dfs(0)
	return ans
}
