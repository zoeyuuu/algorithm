package greedy

import "fmt"

// 860 lemonadeChange easy
// https://leetcode.cn/problems/lemonade-change/

func Problem860() {
	bills := []int{5, 5, 10, 10, 20}
	fmt.Println(lemonadeChange(bills))
}

// 贪心策略：找零的时候尽可能找大面值 这样接下来的操作空间更大
// 本题体现在收到20面值时 优先找10+5 不行再找5*3
func lemonadeChange(bills []int) bool {
	a, b, c := 0, 0, 0
	for i := range bills {
		if bills[i] == 5 {
			a++
		} else if bills[i] == 10 {
			b++
			a--
		} else {
			c++
			if b > 0 {
				b--
				a--
			} else {
				a -= 3
			}
		}
		if a < 0 || b < 0 || c < 0 {
			return false
		}
	}
	return true
}
