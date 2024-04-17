package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int64, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&arr[i])
	// }
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n && scanner.Scan(); i++ {
		num, _ := strconv.Atoi(scanner.Text())
		arr[i] = int64(num)
	}
	var inter [][]int
	start := 0
	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {
			inter = append(inter, []int{start, i})
			start = i
		}
	}
	var result int64
	inter = append(inter, []int{start, n})
	for _, v := range inter {
		result += subarraySum(arr[v[0]:v[1]])
		result %= 10000007
	}
	//fmt.Println(inter)
	fmt.Println(result)
}
func subarraySum(arr []int64) int64 {
	n := len(arr)
	var sum int64
	var ans int64
	pre := make([]int64, n)
	pre[0] = arr[0]
	for i := 1; i < n; i++ {
		pre[i] = (pre[i-1] + arr[i]) % 10000007
	}
	for i := 0; i < n; i++ {
		sum += arr[i]
		for j := i; j < n; j++ {
			if i == 0 {
				ans += pre[j]
			} else {
				ans += pre[j] - pre[i-1]
				ans %= 10000007
			}
		}
	}
	return (ans - sum) % 10000007
}
