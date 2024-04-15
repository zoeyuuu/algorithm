package DP

// 一维
func test_1_wei_bag_problem(weight, value []int, bagWeight int) int {
	// 定义 and 初始化
	dp := make([]int, bagWeight+1)
	// 递推顺序
	for i := 0; i < len(weight); i++ {
		// 这里必须倒序,区别二维,因为二维dp保存了i的状态
		for j := bagWeight; j >= weight[i]; j-- {
			// 递推公式
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	//fmt.Println(dp)
	return dp[bagWeight]
}

// 二维
func test_2_wei_bag_problem1(weight, value []int, bagweight int) int {
	// 定义dp数组  len(weight)*bigweight+1
	dp := make([][]int, len(weight))
	for i, _ := range dp {
		dp[i] = make([]int, bagweight+1)
	}
	// 初始化
	for j := bagweight; j >= weight[0]; j-- {
		dp[0][j] = dp[0][j-weight[0]] + value[0]
	}
	// 递推公式
	for i := 1; i < len(weight); i++ {
		//正序,也可以倒序
		for j := 0; j <= bagweight; j++ {
			if j < weight[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}
	return dp[len(weight)-1][bagweight]
}

func main() {
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	test_2_wei_bag_problem1(weight, value, 4)
}
