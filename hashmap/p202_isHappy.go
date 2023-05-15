package hashmap

import "fmt"

func Problem202() {
	n := 19
	fmt.Println(isHappy(n))
}

func isHappy(n int) bool {
	//哈希表mp存储一系列数字以判断是否进入循环
	mp := map[int]bool{}
	// 终止条件：1、n==1是快乐数 2、
	for ; n != 1 && !mp[n]; n, mp[n] = step(n), true {
	}
	return n == 1
}

// 计算下一个数
func step(n int) int {
	sum := 0
	for n != 0 { //sum每次加上最低位的平方
		sum += (n % 10) * (n % 10)
		n /= 10
	}
	fmt.Println(sum)
	return sum
}
