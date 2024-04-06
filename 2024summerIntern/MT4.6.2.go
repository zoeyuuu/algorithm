package main

/*
小美拿到了一个由复数组成的数组，她想知道其中有多少个实数？
实数：有理数和无理数的总称。其中无理数是无限不循环小数，有理数包括整数和分数。
输入描述：
第一行输入一个正整数，代表数组的大小。
第二行输入n个复数，代表小美拿到的数组。
1\leq n \leq 10^5
后台数据保证复数为a或者a+bi的形式，其中a和b为绝对值不超过10^9的整数。
输出描述：
一个整数，代表实数的数量。
输入示例：
4
-5 5-i 6+3i -4+0i
输出示例：
2
*/

package main

import (
"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]string,n)
	ans := 0
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
		if arr[i][len(arr[i])-1] == 'i'{
			if len(arr[i]) ==1 {
				ans ++
				continue
			}
			j:=len(arr[i])-2
			if arr[i][j] == '0'{
				for ;j>0 && arr[i][j] == '0';j--{}
				if arr[i][j] <= '9' && arr[i][j] >= '1'{
					ans ++
				}
			}else{
				ans ++
			}
		}
	}
	fmt.Println(n-ans)
}
