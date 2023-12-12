package MonotoneStack

import "fmt"

// 2454. 下一个更大元素 IV hard 每日一题
// https://leetcode.cn/problems/next-greater-element-iv/description/?envType=daily-question&envId=2023-12-12

func Problem2454() {
	nums := []int{2, 4, 0, 9, 6}
	fmt.Println(secondGreaterElement1(nums))
}

// 两个单调栈实现：
// p496单调栈实现的时候遇到大于栈顶元素的v 直接弹出栈顶元素并操作 此时v是第一个更大的元素
// 本题要找第二个更大的元素 借鉴p496 弹出的元素原封不动的移动到第二个单调栈t 每次遍历时先判断第二个栈t 同样处理这样比t栈顶大的是第二个更大的元素
func secondGreaterElement1(nums []int) []int {
	res := make([]int, len(nums))
	var s, t []int
	for i := 0; i < len(res); i++ {
		res[i] = -1
	}
	for i, v := range nums {
		// 处理第二个栈t 处理方式与单调栈一模一样
		for len(t) > 0 && v > nums[t[len(t)-1]] {
			res[t[len(t)-1]] = v
			t = t[:len(t)-1]
		}
		// 处理第一个栈s 将栈顶小于v的元素复制到t中
		j := len(s) - 1
		for j >= 0 && v > nums[s[j]] {
			j-- //j--停留在小于等于v的位置
		}
		// 把从 s 弹出的这一整段元素加到 t
		// 先处理t t重所剩元素一定<=v;s[j+1:]>v 所以可以直接append仍然维持t单调栈
		t = append(t, s[j+1:]...)
		// 当前元素（的下标）加到 s 栈顶
		s = append(s[:j+1], i)
	}
	return res
}

// 自己写的 ac不了
// 用一个map存已经遍历过的元素 键：下标 值：更大的元素个数
// 每遍历到元素的时候 比之前大的map 值++ 若>=2 删除该键
// 42 / 51 个通过的测试用例
func secondGreaterElement2(nums []int) []int {
	// 下标：出现次数
	mp := make(map[int]int, len(nums))
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	for i := 0; i < len(nums); i++ {
		mp[i] = 0
		for k, _ := range mp {
			if nums[i] > nums[k] {
				mp[k]++
			}
			if mp[k] >= 2 {
				delete(mp, k)
				res[k] = nums[i]
			}
		}
	}
	return res
}
