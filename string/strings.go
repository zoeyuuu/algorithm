package string

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	a, b, c := strings.Count(str, "y"), strings.Count(str, "o"), strings.Count(str, "u")
	cnt := min(min(a, b), c)
	a, b, c = a-cnt, b-cnt, c-cnt
	ans := strings.Repeat("you", cnt)
	for a > 0 {
		ans += "y"
		a--
	}
	for b > 0 {
		ans += "o"
		b--
	}
	for c > 0 {
		ans += "u"
		c--
	}
	// 剩余加入
	fmt.Println(ans)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
