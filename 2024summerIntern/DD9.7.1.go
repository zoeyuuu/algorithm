package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/*
小C准备参加某个游戏的速通比赛，为此他对该游戏速通了 n次，每次速通记录可以用一个数组 A={a1,a2……am}表示，
其中a表示小C 从游戏开始到第i个游戏节点所花赛的时间，m 为游戏节点的个数。请根据小 C 的速通记录计算出他的理论最佳速通时间，
理论最佳速通时问指:小C在每两个相邻的游戏节点之间所花费的时间均达到了历史最佳记录，在此情况下所花费的总时间。
3 5
1 4 7 9 13
2 3 8 11 14
1 3 7 12 13
*/
func main() {
	var n, m int
	// 使用fmt.Scanf会导致最后的换行符在缓冲区被scanner读取 从而出错
	fmt.Scanln(&n, &m)
	scanner := bufio.NewScanner(os.Stdin)
	path := make([]int, m)
	arr := make([][]int, n)
	for i := range path {
		path[i] = math.MaxInt32
	}
	sum := 0
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		scanner.Scan()
		nums := strings.Fields(scanner.Text())
		for j := 0; j < m; j++ {
			num, _ := strconv.Atoi(nums[j])
			arr[i][j] = num
			if j == 0 {
				path[j] = min(path[j], num)
			} else {
				path[j] = min(path[j], arr[i][j]-arr[i][j-1])
			}
		}
	}
	for i := range path {
		sum += path[i]
	}
	fmt.Println(sum)
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
