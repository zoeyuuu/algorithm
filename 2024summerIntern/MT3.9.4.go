package main

import "fmt"

/*
第四题：小盖的区间删除
小盖拿到了一个大小为n的数组，她希望删除一个区间后，使得剩余所有元素的乘积末尾至少有k个 0。小盖想知道，一共有多少种不同的删除方案？
第一行输入两个正整数n，k。第二行输入n个正整数ai，代表小盖拿到的数组
5 2
2 5 3 4 20
一个整数，代表删除的方案数。
4
第一个方案，删除[3]。第二个方案，删除[4]。第三个方案，删除[3,4]。第四个方案，删除[2]。
*/

// 统计每个数2/5因子个数 0的个数为总量2/5的最小值
// 然后用两层遍历删去区间
// 21/30 组用例通过
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	arr := make([]int, n)
	// 用mp存储数字-因子个数减小复杂度 没提高通过用例个数
	mp2 := make(map[int]int)
	mp5 := make(map[int]int)
	var all2, all5 int
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		mp2[arr[i]] = getFactor2(arr[i])
		mp5[arr[i]] = getFactor5(arr[i])
		all2 += mp2[arr[i]]
		all5 += mp5[arr[i]]
	}
	ans := 0
	for i := 0; i < n; i++ {
		l1, l2 := all2, all5
		for j := i; j < n; j++ {
			l1 -= mp2[arr[j]]
			l2 -= mp5[arr[j]]
			if min(l1, l2) < k {
				break
			}
			ans++
		}
	}
	fmt.Println(ans)
}

// 获得因子3的个数
func getFactor2(a int) int {
	var f2 int
	for a != 0 && a%2 == 0 {
		a /= 2
		f2++
	}
	return f2
}

// 获得因子5的个数
func getFactor5(a int) int {
	var f5 int
	for a != 0 && a%5 == 0 {
		a /= 2
		f5++
	}
	return f5
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
