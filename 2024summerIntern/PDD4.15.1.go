package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var str string
	fmt.Scan(&str)
	dp := make([]int, n)
	dp[2] = abs(int(str[0])-int('P')) + abs(int(str[1])-int('D')) + abs(int(str[2])-int('D'))
	for i := 3; i < n; i++ {
		//fmt.Printf("i=%d,dp[i-1]:%d,", i)
		opt := abs(int(str[i])-int('D')) + abs(int(str[i-1])-int('D')) + abs(int(str[i-2])-int('P'))
		if (i+1)%3 == 0 {
			dp[i] = dp[i-3] + opt
		} else {
			dp[i] = min(dp[i-1], dp[i-3]+opt)
		}
	}
	fmt.Println(dp)
	num := n / 3
	fmt.Printf("%d %d", num, dp[n-1])
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func abs(x int) int {
	if x <= 0 {
		return -x
	}
	return x
}
