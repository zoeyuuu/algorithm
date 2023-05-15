package hashmap

import (
	"fmt"
	"sort"
)

// 49 medium
// 字母异位词分组
// 类似 easy 242

func Problem49() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
	fmt.Println(groupAnagrams2(strs))
}

// 排序
// 时间复杂度O(nklogk) 空间复杂度O(nk) n是字符串数量 k是字符串最大长度
func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{} //键是排序后的字符串 值是字符串切片
	for _, r := range strs {
		str := []byte(r)
		sort.Slice(str, func(i, j int) bool { // 排序
			return str[i] < str[j]
		})
		sortedStr := string(str)
		mp[sortedStr] = append(mp[sortedStr], r)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

// 计数：字符出现次数统计
func groupAnagrams2(strs []string) [][]string {
	mp := map[[26]int][]string{} //键是字符出现次数的统计 值是字符串切片
	for _, v := range strs {
		cnt := [26]int{}
		for _, c := range v {
			cnt[c-'a']++
		}
		mp[cnt] = append(mp[cnt], v)
	}
	ans := make([][]string, 0, len(mp))
	for _, p := range mp {
		ans = append(ans, p)
	}
	return ans
}
