package backtracking

import (
	"fmt"
	"sort"
)

// 这道题目和39.组合总和如下区别：
// 1.本题candidates 中的每个数字在每个组合中只能使用一次    ----->   i+1 dfs(i+1, sum+candidates[i])
// 2.本题数组candidates的元素是有重复的 要去重
// ⭐去重分两种情况
// a)第一种同一层去重 也就是candidates=122222224 target = 7 (本题)  对于第二层会得到无数个124 去重 只需保留一个 124
//  -----> 对数组进行排序 在for循环遍历一层时if i > index && candidates[i] == candidates[i-1] {continue}
// b)第二种是路径(树枝)去重 也就是candidates=122222224 target = 7 不能够得到222的结果
//  -----> 在a)先排序去重基础上加if i > 0 && i == index && candidates[i] == candidates[i-1] {continue} 深度递归的时候判断是否和前一个元素一致
//  -----> 或者直接先排序建立一个新的无重复元素的数组

func Problem40() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum22(candidates, target))
}

func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	n := len(candidates)
	var path []int
	var dfs func(index, sum int)
	sort.Ints(candidates)
	dfs = func(index, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := index; i < n; i++ {
			// 要对同一树层使用过的元素进行跳过 (i > index保证第一次深度遍历到达时不去重）
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			// 改动1：i不用+1
			dfs(i+1, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}

func combinationSum22(candidates []int, target int) [][]int {
	ans := [][]int{}
	n := len(candidates)
	var path []int
	var dfs func(index, sum int)
	sort.Ints(candidates)
	dfs = func(index, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := index; i < n; i++ {
			// 要对同一树层使用过的元素进行跳过 (i > index保证第一次深度遍历到达时不去重）
			if i > index && candidates[i] == candidates[i-1] {
				continue
			}
			if i > 0 && i == index && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			// 改动1：i不用+1
			dfs(i+1, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}
