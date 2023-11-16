package slidingWindow

import "fmt"

// 3. 无重复字符的最长子串 medium
// 比较简单 做出来了
// https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/?envType=study-plan-v2&envId=top-100-liked
// 左右指针维持滑动窗口 用一个map记录当前窗口中出现的字符 键为byte值为int(下标)
// 右指针大循环 当遍历到map中已有的元素时，left指针一直移到mp[s[right]]的下一个位置，同时删除经过的map元素

func Problem3() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
func lengthOfLongestSubstring(s string) int {
	mp := make(map[byte]int)
	sum := 0
	for left, right := 0, 0; right < len(s); right++ {
		// 如果字符已经在窗口中
		if value, ok := mp[s[right]]; ok {
			for ; left <= value; left++ {
				delete(mp, s[left])
			}
			// left = value+1
		}
		// 存 字符-下标
		mp[s[right]] = right
		// 判断当前窗口长短
		if right-left+1 > sum {
			sum = right - left + 1
		}
	}
	return sum
}
