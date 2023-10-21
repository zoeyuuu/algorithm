package doublepointer

import (
	"fmt"
	"sort"
)

func Problem18() {
	nums := []int{1, -2, -5, -4, -3, 3, 3, 5}
	target := -11
	fmt.Println(fourSum1(nums, target))
}
func fourSum1(nums []int, target int) (ans [][]int) {
	n := len(nums)
	if n < 4 {
		return
	}
	sort.Ints(nums)
	// 枚举 a
	for a := 0; a < n; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		// 枚举b
		for b := a + 1; b < n; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}
			// 枚举c
			newTarget := target - nums[a] - nums[b]
			for c, d := b+1, n-1; c < d; c++ {
				if c > b+1 && nums[c] == nums[c-1] {
					continue
				}
				for c < d && nums[c]+nums[d] > newTarget {
					d--
				}
				if c == d {
					break
				}
				if nums[c]+nums[d] == newTarget {
					ans = append(ans, []int{nums[a], nums[b], nums[c], nums[d]})
				}
			}
		}
	}
	return
}
