package main

import (
	"fmt"
	"strings"
)

/*
游游拿到了一个字符串，她想重排这个字符串后，使得该字符串包含尽可能多的"you"连续子串。你能帮帮她吗？
统计you的个数 然后
*/
func main() {
	var str string
	fmt.Scan(&str)
	mp := make(map[byte]int, 3)
	for i := 0; i < len(str); i++ {
		mp[str[i]]++
	}
	a, b, c := mp['y'], mp['o'], mp['u']
	// 另一种统计方式
	// a, b, c := strings.Count(str, "y"), strings.Count(str, "o"), strings.Count(str, "u")
	cnt := min(min(a, b), c)
	a, b, c = a-cnt, b-cnt, c-cnt
	ans := strings.Repeat("you", cnt)
	//for a > 0 { 傻的
	//	ans += "y"
	//	a--
	//}
	ans += strings.Repeat("y", a)
	ans += strings.Repeat("o", b)
	ans += strings.Repeat("u", c)
	// 剩余加入
	fmt.Println(ans)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
