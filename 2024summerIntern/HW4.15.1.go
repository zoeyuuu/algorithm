package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 需要仔细理解题意 首先通过字符串切割读取[][]string类型
// 功能1：去重 根据前三项进行去重 使用mp1[string]bool类型
// 然后使用log[][]string进行存储
// 然后遍历log进行统计 存储在一个客户名-费用的mp3中 统计过程中需要a(存一个mp2 factor-价格的map ; b(过滤异常值
// 最后用一个[][]string拷贝mp3中的数据 使用sort.Slice进行排序
func main() {
	var n int
	fmt.Scan(&n)
	// 为了去重
	mp1 := map[string]bool{}
	var log [][]string
	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		args := strings.Split(str, ",")
		key := args[0] + args[1] + args[2]
		if !mp1[key] {
			mp1[key] = true
			log = append(log, []string{args[1], args[2], args[3]})
		}
	}
	// mp2 factor-单价
	mp2 := map[string]int{}
	var k int
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		var str string
		fmt.Scan(&str)
		args := strings.Split(str, ",")
		num, _ := strconv.Atoi(args[1])
		mp2[args[0]] = num
	}
	mp3 := map[string]int{}
	for _, v := range log {
		num, _ := strconv.Atoi(v[2])
		if num < 0 || num > 100 {
			continue
		}
		fmt.Println(mp2[v[1]], num)
		mp3[v[0]] += mp2[v[1]] * num
	}
	var ans [][]string
	for key, value := range mp3 {
		numStr := strconv.Itoa(value)
		ans = append(ans, []string{key, numStr})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	for _, v := range ans {
		fmt.Printf("%s,%s\n", v[0], v[1])
	}
}

/*
5
1627845600,client1,factorA,10
1627845605,client2,factorB,15
1627845610,client1,factorA,5
1627845615,client1,factorB,8
1627845620,client2,factorB,20
2
factorA,5
factorB,7
*/
