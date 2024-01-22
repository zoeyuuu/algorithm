package interview

// package main
// D:\projects\go_projects\src\algorithm\interview> go run .\hw1-10.go

import (
	"fmt"
	"sort"
)

/*
华为秋招1.10 机试1 简单的任务调度算法
设计一种简单的调度算法，能多根据任务的优先级和执行时长完成调度
1、任务的优先级P值越大，则优先级越高，值起小，则优先级起低，P的取值范图:0<=P <= 20;
2、任务的执行时间T值越大，则执行时间起长，T值为正整数，取值范围:0<T<2^23
3、任务调度需要足如下要求: 1 优先级高的优先度 2.执行时间短的优先调度 3.优先级调度高于执行时间
4.当优先级与执行时间都相同时，编号小的任务优先调度
输入要求：
第一行：任务总数
第二行: 任务列表，以(任务编号，优先级，执行时长)为一个任务
输出要求：
按照调度顺序输出任务给号列表
测试样例1：
输入：
3
1 10 30
2 5 40
3 8 30
输出:1 3 2
测试样例2：
输入：
4
1 10 30
2 5 40
3 8 30
4 5 20
输出：
1 3 4 2

解题思路：
用task []int存储每一个任务的编号、优先级、时常
然后用sort.slice 按照优先级大 时常小 编号小的顺序排序
*/
func main() {
	var n int
	fmt.Scan(&n)
	var tasks [][]int
	for i := 0; i < n; i++ {
		var id, priority, executionTime int
		fmt.Scan(&id, &priority, &executionTime)
		task := []int{id, priority, executionTime}
		tasks = append(tasks, task)
	}
	sort.Slice(tasks, func(i, j int) bool {
		// 优先级降序
		if tasks[i][1] != tasks[j][1] {
			return tasks[i][1] > tasks[j][1]
		}
		// 时常升序
		if tasks[i][2] != tasks[j][2] {
			return tasks[i][2] < tasks[j][2]
		}
		// 编号升序
		return tasks[i][0] < tasks[j][0]
	})
	for i, task := range tasks {
		fmt.Print(task[0])
		if i < len(tasks) {
			fmt.Print(" ")
		}
	}
}
