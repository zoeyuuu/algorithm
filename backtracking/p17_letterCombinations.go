package backtracking

// 17. 电话号码的字母组合 medium
// 下一次用回溯算法模板再写一遍

import "fmt"

func Problem17() {
	var digits string = "2"
	fmt.Println(letterCombination(digits))
}

var mp map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombination(digits string) (ans []string) {
	n := len(digits)
	if n == 0 {
		return
	}
	path := make([]byte, n)
	var dfs func(i int) //函数变量
	dfs = func(i int) {
		if i == n { //最后一个数字确定
			ans = append(ans, string(path))
			return
		}
		digit := string(digits[i])
		letters := mp[digit]
		for k := 0; k < len(letters); k++ {
			path[i] = letters[k]
			dfs(i + 1)
		}
	}
	dfs(0)
	return
}
