package interview

// 根据输入条件 用一个数组存储所有出现的数
// 然后遍历一遍 输出

import (
	"fmt"
)

func main() {
	N := 0
	fmt.Scan(&N)
	narray := make([]int, 501)
	n := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&n)
		narray[n] = 1
	}
	for i, _ := range narray {
		if narray[i] > 0 {
			fmt.Println(i)
		}
	}
}
