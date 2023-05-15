package array

// easy
import "fmt"

func problem2373() {
	grid := [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}
	fmt.Println(largestLocal(grid))
}

// 两个for循环进行遍历 每次在一个3*3的
func largestLocal(grid [][]int) [][]int {
	n := len(grid)
	ans := make([][]int, n-2)
	for i := 0; i < n-2; i++ {
		ans[i] = make([]int, n-2) //注意这里是"=" 而不是 ":="
		for j := 0; j < n-2; j++ {
			temp := grid[i][j]
			for k := i; k < i+3; k++ {
				for l := j; l < j+3; l++ { // k < i+3 而不是 k < k+3
					if grid[k][l] > temp {
						temp = grid[k][l]
					}
				}
			}
			ans[i][j] = temp
		}
	}
	return ans
}
