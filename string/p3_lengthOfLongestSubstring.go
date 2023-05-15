package string

import "fmt"

// 3.无重复字符的最长子串 medium
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

/* 算法思路
维持一个无重复字符的滑动窗口，利用双指针
left right初始都为0 如果字符一直不重复 right++ 并存入哈希表
字符的重复利用动态哈希表存储来维持
遇到重复的字符left++ 并移除哈希表对应元素 直到哈希表中没有该元素
*/

func Problem3() {
	var a = ""
	fmt.Println(lengthOfLongestSubstring(a))
}

/*特殊情况最好是单独写
  if n <= 1{
	return n
  }
  n=0时right<n不成立不进入循环，n=1时也成立
*/

func lengthOfLongestSubstring(str string) int {
	mp := map[byte]bool{}
	n := len(str)
	fmt.Println(n)
	ans := 0
	for left, right := 0, 0; right < n; right++ {
		for mp[str[right]] { //for的判断只能是bool类型 mp的值定义为bool方便方便这里的判断
			delete(mp, str[left])
			left++
		}
		mp[str[right]] = true //这里相当于给mp[str[0]]赋值true
		ans = max(ans, right-left+1)
	}
	return ans
}
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
