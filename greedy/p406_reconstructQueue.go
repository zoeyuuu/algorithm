package greedy

import (
	"container/list"
	"fmt"
	"sort"
)

// 406 reconstructQueue medium 能到hard
// https://leetcode.cn/problems/queue-reconstruction-by-height/
// 思路：按照身高由大到小进行排序 然后从数组头到尾按照people[i][1](排在前面大于等于自己的人数)进行插入构造新数组
// 理解：按照身高排序进行插入排序 那后面插入的元素更小绝对不可能构成逆序 所以在插入该元素的时候满足逆序个数people[i][1]即可

func Problem406() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
}

// 方法一 数组要移动不少 已经可以了
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			//身高相同的时候 按照K值从小到大排序（如果K值大的先插入 那身高相同K值小的进行插入的时候会导致出错前者的逆序+1）
			return people[i][1] > people[j][1]
		}
		return people[i][0] > people[j][0] //按照身高从高到底排序
	})
	for i, p := range people {
		// p[1]代表大于等于的元素个数 即插入的下标
		// 将下标p[1]~i-1复制到p[1]+1~i
		copy(people[p[1]+1:i+1], people[p[1]:i])
		people[p[1]] = p //在p[1]下标插入p
	}
	return people
}

// 方法二 链表实现
func reconstructQueue1(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1] //当身高相同时，将K按照从小到大排序
		}
		return people[i][0] > people[j][0]
	})
	l := list.New() //创建链表
	for i := 0; i < len(people); i++ {
		position := people[i][1]      //position的值代表前面有几个元素
		mark := l.PushBack(people[i]) //插入元素
		e := l.Front()
		for position != 0 { //获取相对位置
			position--
			e = e.Next()
		}
		l.MoveBefore(mark, e) //移动位置

	}
	res := [][]int{}
	for e := l.Front(); e != nil; e = e.Next() {
		res = append(res, e.Value.([]int))
	}
	return res
}
