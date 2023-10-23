package DP

import "fmt"

// 139. 单词拆分 meidum
// https://leetcode.cn/problems/word-break/description/
// dp[i] 0~i-1的字符串能拆分
// 递推式：dp[i]= dp[j]&&check(j,i-1)   0~i-1的字符串能拆分:j遍历位置0~i-1,dp[j]能拆分(0~j-1)并且j~i-1在字符串字典中
// 初始化 dp[0] = true
// 也可以看作完全背包 但是要求遍历顺序先是背包容量，后是物品 因为这里是bool类型递推式 与遍历顺序有关系 考虑一下反例
// s = "applepenapple", wordDict = ["pen","apple"] 如果先遍历物品pen，此时dp[5]=flase(apple还未遍历),dp[8]=false&&true==flase
// 所以本题一定要遍历容量 后物品 bool类型递推式

func Problem139() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
}
func wordBreak(s string, wordDict []string) bool {
	mp := make(map[string]bool)
	// 子串字典
	for _, v := range wordDict {
		mp[v] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i <= n; i++ {
		// j < i <=都行 j=i没有意义
		for j := 0; j < i; j++ {
			// 0~i-1为true：0~j-1可拆分&&j~i-1在字典内
			if dp[j] && mp[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	// 0~n-1可拆分
	return dp[n]
}
