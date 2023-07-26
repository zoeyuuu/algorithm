package interview

// https://www.nowcoder.com/practice/fe298c55694f4ed39e256170ff2c205f?tpId=37&tqId=21245&ru=/exam/oj
// 三个空汽水瓶可以换一瓶汽水，允许向老板借空汽水瓶（但是必须要归还）。
// 小张手上有n个空汽水瓶，她想知道自己最多可以喝到多少瓶汽水。
// 两个空瓶可以借一个空瓶换一瓶
import "fmt"

func Bottle() {
	var input int
	for {
		fmt.Scan(&input)

		if input == 0 {
			break
		}
		fmt.Println(f(input))
	}
}

// 递归 f(i-2) 加两个空瓶可以换一个
func f(i int) int {
	switch i {
	case 1:
		return 0
	case 2:
		return 1
	default:
		return f(i-2) + 1
	}
}

// 两个空瓶可以借一个空瓶换一瓶
func f2(i int) int {
	return i / 2
}
