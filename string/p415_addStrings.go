package string

import "strconv"

// 415 字符串相加 easy
// 模拟数字加法 从最后一位开始
func addStrings(num1 string, num2 string) string {
	add := 0
	var ans string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add > 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		sum := x + y + add
		add = sum / 10
		ans = strconv.Itoa(sum%10) + ans
	}
	return ans
}
