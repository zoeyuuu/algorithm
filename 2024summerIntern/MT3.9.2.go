package main

import (
	"fmt"
)

// 27/30 组用例通过
func main() {
	var n, q int
	fmt.Scan(&n, &q)
	arr := make([]int, n)
	sum := 0
	num := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
		sum += arr[i]
		if arr[i] == 0 {
			num++
		}
	}
	var l, r int
	for i := 0; i < q; i++ {
		fmt.Scan(&l, &r)
		//fmt.Printf("%d %d\n", sum+l, sum+r)
		fmt.Println(sum+num*l, sum+num*r)
	}
}
