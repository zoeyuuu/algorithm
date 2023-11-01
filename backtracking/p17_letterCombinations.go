package backtracking

// 17. 电话号码的字母组合 medium
// 下一次用回溯算法模板再写一遍

import "fmt"

func Problem17() {
	var digits string = "23"
	fmt.Println(letterCombination(digits))
}

// 终止条件是i和n进行比较
func letterCombination(digits string) (ans []string) {
	mp := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	n := len(digits)
	if n == 0 {
		return
	}
	path := make([]byte, n)
	var dfs func(i int)
	dfs = func(i int) {
		//在下一层写递归终止条件
		if i == n {
			ans = append(ans, string(path))
			return
		}
		// byte转成string类型 查找map的键
		digit := string(digits[i])
		letters := mp[digit]
		for k := 0; k < len(letters); k++ {
			path[i] = letters[k]
			dfs(i + 1)
		}
	}
	// i是下标
	dfs(0)
	return
}
