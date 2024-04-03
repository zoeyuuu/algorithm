package main

import (
	"fmt"
)

func main() {
	var n, k int
	var str string
	fmt.Scan(&n, &k)
	fmt.Scan(&str)
	sum := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'M' || str[i] == 'T' {
			sum++
		}
	}
	result := min(sum+k, len(str))
	fmt.Println(result)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
