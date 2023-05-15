package hashmap

// 242easy
// 有效的字母异位词 题目要求只有小写字母

import (
	"fmt"
	"sort"
)

func Problem242() {
	s1 := "abc中午"
	s2 := "bac中午"
	fmt.Println(isAnagram_sort(s1, s2))
	//fmt.Println(isAnagram_hashmap(s1, s2))
	fmt.Println(isAnagram_hashmap2(s1, s2))
}

// 字符串转换成字节切片rune类型切片排序进行比较
func isAnagram_sort(s, t string) bool {
	// 字符串（不可变）转换为字节序列进行排序比较等 ASCII码
	// s1, s2 := []byte(s), []byte(t)
	// 上一行代码不能处理unicode字符 例如中文字符转换后占3个[]byte字节
	// 排序后string(s1)虽然能比较但会出错乱码

	// []rune类型转换为整数切片 可以处理unicode字符 （例如中文字符）
	s1, s2 := []rune(s), []rune(t)
	fmt.Println(s1, s2)
	sort.Slice(s1, func(i, j int) bool { //基于快排 O(nlogn)
		return s1[i] < s1[j] //如果是true就交换
	})
	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	return string(s1) == string(s2)
}

// 26数组记录 哈希表  只能处理小写字符
func isAnagram_hashmap(s, t string) bool {
	record := [100000]int{}
	for _, r := range s {
		record[r-'a']++
	}
	for _, r := range t {
		record[r-'a']--
	}
	return record == [100000]int{}
}

func isAnagram_hashmap2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnd := map[rune]int{} //为了处理unicode字符 使用哈希表
	for _, r := range s {
		cnd[r]++
	}
	for _, r := range t {
		cnd[r]--
		if cnd[r] < 0 {
			return false
		}
	}
	return true
}
