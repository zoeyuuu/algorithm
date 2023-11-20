package backtracking

// 37. 解数独 hard
// 回溯法 二维遍历
// 先判断是否能填 然后尝试1~9 都不行则返回false 二维都遍历完那么返回true
func solveSudoku(board [][]byte) {
	var backtracking func(board [][]byte) bool
	backtracking = func(board [][]byte) bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				//判断此位置是否适合填数字
				if board[i][j] != '.' {
					continue
				}
				//尝试填1-9
				for k := '1'; k <= '9'; k++ {
					if isvalid(i, j, byte(k), board) == true { //如果满足要求就填
						board[i][j] = byte(k)
						if backtracking(board) == true {
							return true
						}
						board[i][j] = '.'
					}
				}
				// 9个数都试完了，都不行，那么就返回false
				return false
			}
		}
		// 遍历完没有返回false，说明找到了合适棋盘位置了
		return true
	}
	backtracking(board)
}

// 判断填入数字是否满足要求
func isvalid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { //行
		if board[row][i] == k {
			return false
		}
	}
	for i := 0; i < 9; i++ { //列
		if board[i][col] == k {
			return false
		}
	}
	//方格
	startrow := (row / 3) * 3
	startcol := (col / 3) * 3
	for i := startrow; i < startrow+3; i++ {
		for j := startcol; j < startcol+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}
	return true
}
