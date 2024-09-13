package main

import "fmt"

/*
一家木材厂需要加工三根圆木。这三根圆木长度分别为 a,b,c，一共需要进行不超过n次加工程序。第i道加工程序需要选择其中一根长度严格大于i的圆木，将其切割，使其长度减少i。被切下的部分不再进入后续的加工流程。如果这三根圆木的长度能够组成一个面积大于0的三角形，那么就称此时的圆木长度三元组(ab,c)是好的 现在的问题是:一共可能形成多少种好的三元组?
输入描述
输入仅一行四个正整数n,a,b,c。对于100%的数据，1<=n,a,b,c<=100
输出描述
输出一行，一个整数，表示好的三元组的个数
样例输入
5 3 4 5
样例输出
10
*/

// 自测可以运行 但是应该可以剪枝条件
func main() {
	var n, a, b, c int
	fmt.Scanln(&n, &a, &b, &c)
	var ans [][]int
	var dfs func(k int)
	dfs = func(k int) {
		if isTriangle(a, b, c) {
			ans = append(ans, []int{a, b, c})
		}
		if k > n {
			return
		}
		if a-k > 0 {
			a -= k
			dfs(k + 1)
			a += k
		}
		if b-k > 0 {
			b -= k
			dfs(k + 1)
			b += k
		}
		if c-k > 0 {
			c -= k
			dfs(k + 1)
			c += k
		}
	}
	dfs(1)
	fmt.Println(ans)
	fmt.Println(len(ans))
}
func isTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}
