package slidingWindow

import "fmt"

// 438. 找到字符串中所有字母异位词 medium
// 没做出来 当时思路：
// 存储一个匹配串字符-出现次数的map 遍历的时候遇到字符map中相应出现次数--(消耗) 但想利用len(mp)==0 作为匹配完成的终止条件
// 涉及一些点1.出现次数为1的时候再减就要delete键值对 2.遇到map不存在的字符要分为匹配串中存在与否 是否重新加入map比较麻烦

func Problem438() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams1(s, p))
	fmt.Println(findAnagrams2(s, p))
}

// 方法一：改进之前的想法 利用right-left+1 =len(p)作为终止条件 不需要考虑删除键的问题
// 遇到出现次数=0的情况但仍需-1(不够小消耗) left指针右移
// 时间复杂度：O(m+n)
// 空间复杂度：O(Σ)
func findAnagrams1(s string, p string) (ans []int) {
	m, n := len(s), len(p)
	if m < n {
		return
	}
	// 存储的是需要消耗的数量
	mp := make(map[byte]int)
	for i := range p {
		mp[p[i]]++
	}
	for l, r := 0, 0; r < m; r++ {
		// 先将mp[s[r]]--
		mp[s[r]]--
		// 一旦mp[s[r]]<0 左指针右移 消耗的还回去
		for mp[s[r]] < 0 {
			mp[s[l]]++
			l++
		}
		// 长度为n的时候找到结果
		if r-l+1 == n {
			ans = append(ans, l)
		}
	}
	return ans
}

// 方法二：滑动窗口大小不变直接右移
// golang存储成大小为26数组可以直接比较大小 (s 和 p 仅包含小写字母)
// sCount,pCount直接进行比较
// 时间复杂度：O(n+(m−n)×Σ)，其中 m 为字符串 s的长度，n 为字符串 p 的长度，Σ为所有可能的字符数。
// 空间复杂度：O(Σ)
func findAnagrams2(s string, p string) (ans []int) {
	m, n := len(s), len(p)
	if m < n {
		return
	}
	// 数组实现哈希表
	var sCount, pCount [26]int
	for i := range p {
		sCount[s[i]-'a']++
		pCount[p[i]-'a']++
	}
	// 很容易错 由于for循环中要进行s[i+n]这种判断 很容易越界
	// ans = append(ans, i+1)在移动之后判断 i要+1
	// 这样一开始的时候 要提前判断=0的情况
	if sCount == pCount {
		ans = append(ans, 0)
	}
	// 边界条件(i<=m-n) 自己算一下
	for i := 0; i < m-n; i++ {
		sCount[s[i]-'a']--
		sCount[s[i+n]-'a']++
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return ans
}
