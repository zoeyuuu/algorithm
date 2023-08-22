package greedy

// 738. 单调递增的数字 medium
// https://leetcode.cn/problems/monotone-increasing-digits/description/
// 方法一
// 自己思路：1.找到第一个逆序数的位置 2.让此位置的数-- 后边的数都变成‘9’
//          3.由于-- 要保证递增 所以要保证该位置大于上一个位置 不能等于
// 注意点：1.整型转换成string类型再转换成byte类型便于更改 2.循环的时候涉及比较i,i+1下标注意边界防止数组越界
//        3.找到需要改变的位置对于for循环要break 否则会出错

// 方法二 **：代码随想录
// 为了防止自减改变递增顺序 改为从后向前遍历 这样自减可能的带来的逆序会在接下来的遍历的时候消除
// 代码更直观 但复杂度可能高一点 因为一遇到逆序就要一个for循环改变

import (
	"fmt"
	"strconv"
)

func Problem738() {
	n := 127768
	fmt.Println(monotoneIncreasingDigits1(n))
	fmt.Println(monotoneIncreasingDigits2(n))
}

func monotoneIncreasingDigits1(n int) int {
	// int -> string -> []byte 便于更改
	s := strconv.Itoa(n)
	ss := []byte(s)
	// i:0~i-2 因为循环中有ss[i+1]
	for i := 0; i < len(ss)-1; i++ {
		// 如果有逆序
		if ss[i] > ss[i+1] {
			flag := 0
			// 找到能够--的位置flag 并且保持递增
			// 循环j:1~i
			for j := i; j > 0; j-- {
				if ss[j-1] < ss[j] {
					flag = j
					//处理完break
					break
				}
			}
			ss[flag]--
			for j := flag + 1; j < len(ss); j++ {
				ss[j] = '9'
			}
			//处理完break
			break
		}
		//没有逆序什么操作也不做
	}
	// []byte -> string -> int
	res, _ := strconv.Atoi(string(ss))
	return res
}

func monotoneIncreasingDigits2(N int) int {
	s := strconv.Itoa(N) //将数字转为字符串，方便使用下标
	ss := []byte(s)      //将字符串转为byte数组，方便更改。
	n := len(ss)
	if n <= 1 {
		return N
	}
	for i := n - 1; i > 0; i-- {
		if ss[i-1] > ss[i] { //前一个大于后一位,前一位减1，后面的全部置为9
			ss[i-1] -= 1
			for j := i; j < n; j++ { //后面的全部置为9
				ss[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(ss))
	return res
}
