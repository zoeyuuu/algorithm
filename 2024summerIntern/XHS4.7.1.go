package main

import (
	"fmt"
	"sort"
)

// 根据点赞数，收藏数指标排序
// 最后输出
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	var arr [][]int
	for i := 0; i < n; i++ {
		var like, collect int
		fmt.Scan(&like, &collect)
		// 编号要+1
		arr = append(arr, []int{i + 1, like, collect, like + collect*2})
	}
	sort.Slice(arr, func(i, j int) bool {
		// 先按 支持度降序
		if arr[i][3] != arr[j][3] {
			return arr[i][3] > arr[j][3]
		}
		// 支持度相同 收藏降序
		if arr[i][2] != arr[j][2] {
			return arr[i][2] > arr[j][2]
		}
		// 都相同选择编号更小的
		return arr[i][0] < arr[j][0]
	})
	//fmt.Println(arr)
	// 初始化长度要设置成0
	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		//fmt.Println(arr[i][0])
		ans = append(ans, arr[i][0])
	}
	//fmt.Println(ans)
	sort.Ints(ans)
	for i := 0; i < k; i++ {
		fmt.Print(ans[i])
		if i < k-1 {
			fmt.Print(" ")
		}
	}
}

/* https://mp.weixin.qq.com/s/N5I7eEHsS9DUIR49uvvawQ
小苯是“小红书app”的一名博主,这天他想要给自己的“铁粉”送一些礼物。他有n名粉丝，编号从1到n,但他只能选择其中k名送礼物,他决定选择其中对他支持力度最大的前k名粉丝。
(如果两名支持力度相同,则优先选择收藏数更多的,如果都一样,则优先选择编号更小的(因为这意味着他关注小苯的时间更早))
具体的:每名粉丝如果每给小苯点一次赞,则他对小苯就增加了1点支持力度,如果他每收藏小苯的一篇文章,则他对小苯增加2点支持力度。
现在小苯想知道,他应该选择哪k名粉丝送出礼物,请你帮帮他吧。
输入描述：
输入包括n+1行
第一行两个正整数n, k(,分别表示对小苯有过支持的粉丝个数，以及小苯选择送礼的粉丝个数。
接下来n行,每行两个整数x, y(),表示第i位粉丝给小苯点过×次赞,收藏过y个小苯的文章,
输出描述：
输出包含一行k个正整数,表示小苯选择出送礼物的粉丝们的编号。(按照升序输出)
输入：
4 2
1 2
2 1
3 0
1 3
输出：
1 4
*/
