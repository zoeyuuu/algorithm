package main

import (
	"fmt"
	"sort"
)

// 自定义字符顺序排序

func main() {
	// 定义字符映射
	mp := map[byte]int{
		'a': 1, 'b': 2, 'c': 3, 'd': 4, // 可以继续添加映射
	}

	// 字符串切片
	strList := []string{"dab", "abc", "cab"}

	// 自定义排序逻辑
	sort.Slice(strList, func(i, j int) bool {
		// 获取两个字符串
		str1 := strList[i]
		str2 := strList[j]

		// 按字符顺序比较
		minLen := len(str1)
		if len(str2) < minLen {
			minLen = len(str2)
		}

		for k := 0; k < minLen; k++ {
			// 获取每个字符的映射值
			val1 := mp[str1[k]]
			val2 := mp[str2[k]]

			// 如果两个字符的值不相等，按映射值排序
			if val1 != val2 {
				return val1 < val2
			}
		}

		// 如果字符前缀相同，按字符串长度排序
		return len(str1) < len(str2)
	})

	// 打印排序后的结果
	fmt.Println(strList)
}
