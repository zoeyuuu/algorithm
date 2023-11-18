package sort

import (
	"fmt"
	"math"
	"sort"
)

func Main() {
	// 1.整型
	nums := []int{2, 4, 3, 5, 5, 7}
	// 将nums类型转换为sort.IntSlice的sort.Interface接口
	sort.Sort(sort.IntSlice(nums))
	// sort.Ints自动将[]int类型转换为sort.IntSlice类型 然后调用sort.Sort()
	sort.Ints(nums)                    // [2 3 4 5 5 7]
	sort.IsSorted(sort.IntSlice(nums)) // true
	sort.IntsAreSorted(nums)           //true
	// search
	// 在已排序的整数切片中查找某个值，并返回该值在切片中的索引;如果v不在切片中，index将是v应该插入的位置。
	index1 := sort.SearchInts(nums, 5)
	index2 := sort.Search(len(nums), func(i int) bool { return nums[i] >= 5 }) // 等价上一条(匿名函数希望查找大于等于 6 的元素)
	index3 := sort.Search(len(nums), func(i int) bool { return nums[i] > 5 })  // 匿名函数希望查找大于6的元素
	fmt.Println(index1, index2, index3)                                        // 3 3 5

	// 2.浮点型
	nums1 := []float64{0.0, 1.2, 0.33, math.NaN()}
	// sort.Sort(sort.Float64Slice(nums1))
	sort.Float64s(nums1)                    // [NaN 0 0.33 1.2]
	sort.IsSorted(sort.Float64Slice(nums1)) //true
	sort.Float64sAreSorted(nums1)           //true
	// search
	index1 = sort.SearchFloat64s(nums1, 0.33)
	index2 = sort.Search(len(nums1), func(i int) bool { return nums1[i] >= 0.33 })
	index3 = sort.Search(len(nums1), func(i int) bool { return nums1[i] > 0.33 })
	fmt.Println(index1, index2, index3) // 2 2 3

	// 3.字符串型
	// ASCII 编码: 0~9:48~57  A:65 a:97
	nums2 := []string{"abc", "aba", "Jor", "123", "Jo"}
	// sort.Sort(sort.StringSlice(nums2))
	sort.Strings(nums2)                    // [123 Jo Jor aba abc]
	sort.IsSorted(sort.StringSlice(nums2)) //true
	sort.StringsAreSorted(nums2)           //true
	index1 = sort.SearchStrings(nums2, "Jo")
	index2 = sort.Search(len(nums2), func(i int) bool { return nums2[i] >= "Jo" })
	fmt.Println(index1, index2) // 1 1

	// 4. 降序排序
	nums3 := []int{2, 31, 5, 6, 3}
	// sort.Reverse 返回的类型是一个新的匿名结构体,实现了sort.Interface接口,但覆盖了Less方法,变成逆序排序 -> Less(j, i)
	// 需要最后连sort.Sort()
	sort.Sort(sort.Reverse(sort.IntSlice(nums3)))
	fmt.Println(nums3) //[31 6 5 3 2]

	// 5.sort.Slice
	// 会使用提供的Less函数对整数切片nums进行排序
	// less函数 < 表示 表示只有在 nums[i] 小于 nums[j] 时，它们的顺序才会被调整 排序是稳定的
	// <= 排序不稳定
	nums4 := []string{"aac", "aba", "Kor", "123", "Jo"}
	sort.Slice(nums4, func(i, j int) bool {
		return nums4[i][1] < nums4[j][1]
	}) // 按照字符串的二个字符排序
	fmt.Println(nums4) //[123 aac aba Jo Kor]
	nums5 := []int{2, 4, 3, 5, 5, 7}
	sort.Slice(nums5, func(i, j int) bool { return nums5[i] < nums5[j] }) //[2 3 4 5 5 7]
	sort.Slice(nums5, func(i, j int) bool { return nums5[i] > nums5[j] }) //[7 5 5 4 3 2]

	// 6.func Stable(data Interface) 保持相同元素相对位置不变 稳定排序
	nums = []int{2, 4, 3, 8, 5, 7}
	sort.Stable(sort.IntSlice(nums))
	sort.Slice(nums, func(i, j int) bool { return nums5[i] < nums5[j] }) //等效
}
