package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{1, 2, 5, 3, 4}
	nums3 := []int{1, 5, 4, 2, 3}
	fmt.Println(checkArray(nums1))
	fmt.Println(checkArray(nums2))
	fmt.Println(checkArray(nums3))
}
func checkArray(arr []int) int {
	n := len(arr)
	flag := 0
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			continue
		} else {
			flag++
			if flag > 1 {
				return 1
			}
			if arr[i] > arr[i-2] {
				continue
			} else {
				return 1
			}
		}
	}
	return 3 - flag
}
