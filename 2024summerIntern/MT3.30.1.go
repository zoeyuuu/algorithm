package main

/*
众所周知，“全国大学生英语四级考试”（以下简称CET-4）的满分为 710 分。
在经过三次的 CET-4 考试后，小美终于如愿通过了四级。
已知 CET-4 的总分数由三大项构成，分别是：听力，阅读，写作。
现在已知小美的总分为 k，且他的“写作”得分比“听力”高了 x 分，比“阅读”得分低了 y 分，你能求出他的三个大项分别得了多少分吗？
*/
import (
	"fmt"
)

func main() {
	var k, x, y int
	fmt.Scan(&k, &x, &y)
	var a, b, c int
	b = (k + 2*y + x) / 3
	a = (k - 2*x - y) / 3
	c = k - a - b
	fmt.Printf("%d %d %d", a, b, c)
}
