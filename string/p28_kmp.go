package string

import "fmt"

// 28. 找出字符串中第一个匹配项的下标 easy
// KMP: medium ~ hard
// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/

func Problem28() {
	haystack := "esadbutsad"
	needle := "sad"
	fmt.Println(strStr(haystack, needle))
}
func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	next := make([]int, n)
	// 求next数组
	getNext(next, needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		// 匹配不成功时 j回退到上一个next[j-1]的位置
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		// j==n 全部匹配完 返回起始下标
		if j == n {
			return i - j + 1
		}
	}
	return -1
}

// kmp
// 求next数组 next[j]:0~j下标的子串最长相同前后缀的长度 也是下一个继续比较的位置 index+1
// （主串模式串当遇到不匹配时 找到前一个位置的next数组值回退）
// 求next数组 时间复杂度O(M)，M为模式串的长度
// next数组初始化：next[0]=0 只有一个字母的串前后缀都为空 （前缀是指去除尾部元素的子串）
// i=1 外层循环(1~n-1) 求i处的next值
// j=0 i表示的是后缀末尾，j表示前缀末尾
// 每一次循环前 next[0~i-1]都是已经求好的next数组 **并且0~j-1前缀 和x~i-1都是匹配的**
// 情况1. s[i]==s[j]   0~j和x~i匹配 然后j++ (i++外循环）  next[i]=j  j先++在赋值 next[j]表示的是下一个比较的位置
// 情况2. s[i]!=s[j] 与kmp思路一致 利用已经比较过的部分 j指针回退 j=next[j-1] 继续比较直至s[i]=s[j]或j到0
// next[j-1]=k 代表s[0~k-1]==s[x~j-1] == s[x~i-1] 此时j回退至next[j-1]=k 比较s[k]和s[i]
func getNext(next []int, s string) {
	next[0] = 0
	j := 0
	for i := 1; i < len(s); i++ {
		// 不相等回退
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++ // j先++
		}
		next[i] = j
	}
}
