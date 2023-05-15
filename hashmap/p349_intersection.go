package hashmap

// 判断两个数组的交集
//

import "fmt"

func Problem349() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
func intersection(nums1, nums2 []int) []int {
	var ans []int
	map1 := map[int]struct{}{}
	for _, v := range nums1 {
		map1[v] = struct{}{}
	}
	map2 := map[int]struct{}{}
	for _, v := range nums2 {
		map2[v] = struct{}{}
	}
	if len(map1) > len(map2) { // 为了最后遍历短的那个
		map1, map2 = map2, map1
	}
	for v := range map1 {
		if _, ok := map2[v]; ok { // 判断map中键是否存在
			ans = append(ans, v)
		}
	}
	return ans
}
