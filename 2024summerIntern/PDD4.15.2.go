package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var inter [][]int
	start := 0
	for i := 1; i < n; i++ {
		if arr[i] == arr[i-1] {
			inter = append(inter, arr[start:i])
			start = i
		}
	}
	result := 0
	inter = append(inter, arr[start:])
	for _, v := range inter {
		result += subarraySum(v)
	}
	//fmt.Println(inter)
	fmt.Println(result)
}
func subarraySum(arr []int) int {
	n := len(arr)
	sum := 0
	ans := 0
	pre := make([]int, n)
	pre[0] = arr[0]
	for i := 1; i < n; i++ {
		pre[i] = pre[i-1] + arr[i]
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
	return ans - sum
}
