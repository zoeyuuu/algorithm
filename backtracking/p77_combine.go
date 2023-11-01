package backtracking

// 77 组合 medium
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
// 回溯法

import "fmt"

func Problem77() {
	//var n, k int
	//_, err := fmt.Scan(&n, &k)
	//if err != nil {
	//	return
	//}
	fmt.Println(combine1(4, 2))
	//fmt.Println(combine2(n, k))
}

// 23.11.1 用内部函数的方法 ⭐
func combine1(n, k int) [][]int {
	var res [][]int
	// path初始化不能用path:=make([]int,k) 否则一开始就有k的长度直接退出
	path1 := []int{}
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if len(path1) == k {
			temp := make([]int, k)
			copy(temp, path1)
			res = append(res, temp)
			// 终止条件根据path长度 返回之后由上一层进行回溯操作
			return
		}
		// 剪枝：可选取的元素个数小于还需要的元素个数就停止
		// n-x+1=k-len(path1) 或者用n-i+1 < k-len(path) -> break
		for i := startIndex; i <= n-k+len(path1)+1; i++ {
			path1 = append(path1, i)
			dfs(i + 1)
			// 回溯
			path1 = path1[:len(path1)-1]
		}
	}
	dfs(1)
	return res
}

// 23.5.15 用全局变量写的
//var (
//	path []int
//	ans  [][]int
//)
//func combine2(n, k int) [][]int {
//	// 这里数组元素即是1~n 可以直接用数字代替start下标
//	dfs77(n, k, 1)
//	return ans
//}
//func dfs77(n, k, start int) {
//	if len(path) == k {
//		// 需要创建一个temp再append 否则后续对path的修改会影响
//		tmp := make([]int, k)
//		copy(tmp, path)
//		ans = append(ans, tmp)
//		return
//	}
//	for i := start; i <= n; i++ {
//		// 剪枝：可选取的元素个数小于还需要的元素个数就停止
//		if n-i+1 < k-len(path) {
//			break
//		}
//		path = append(path, i)
//		dfs77(n, k, i+1)
//		path = path[:len(path)-1]
//	}
//}
