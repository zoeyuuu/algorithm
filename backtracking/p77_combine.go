package backtracking

// 77 组合 medium
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 看代码随想录讲解：https://programmercarl.com/0077.%E7%BB%84%E5%90%88.html#%E5%9B%9E%E6%BA%AF%E6%B3%95%E4%B8%89%E9%83%A8%E6%9B%B2
// 回溯法的搜索过程就是一个树型结构的遍历过程，在如下图中，可以看出for循环用来横向遍历，递归的过程是纵向遍历。

import "fmt"

var (
	path []int
	ans  [][]int
)

func Problem77() {
	var n, k int
	_, err := fmt.Scan(&n, &k)
	if err != nil {
		return
	}
	fmt.Println(combine(n, k))
}

func combine(n, k int) [][]int {
	dfs77(n, k, 1) //
	return ans
}
func dfs77(n, k, start int) {
	if len(path) == k {
		// 需要创建一个temp再append 否则后续对path的修改会影响
		tmp := make([]int, k)
		copy(tmp, path)
		ans = append(ans, tmp)
		return
	}
	for i := start; i <= n; i++ {
		path = append(path, i)
		//一开始错写成了 dfs77(n, k, start+1)
		dfs77(n, k, i+1)
		path = path[:len(path)-1]
	}
}
