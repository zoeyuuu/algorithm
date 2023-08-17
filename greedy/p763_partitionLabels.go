package greedy

import "fmt"

// 763. 划分字母区间 medium
// https://leetcode.cn/problems/partition-labels/description/
// 首先读懂意思 **：对于划分出的区间 每一个字母只在该区间出现 且区间尽可能小
// 对于划分区间 每个元素最后一次出现的位置都<=区间尾部
// 第一遍遍历记录每一个字母出现的最后位置
// 第二遍遍历 保证end之前所有字母的最后位置都<=end 否则更新end 当下标i==end时得到一个区间 start=end+1

func Problem763() {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}
func partitionLabels(s string) []int {
	hashmap := make(map[rune]int)
	for i, str := range s {
		hashmap[str] = i
	}
	partition := []int{}
	//start,end记录字符串首尾位置
	start, end := 0, 0
	for i, str := range s {
		if hashmap[str] > end {
			end = hashmap[str]
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return partition
}
