/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	// ITERATION 1:
	// rowSet := make(map[byte]bool)
	// colSet := make(map[byte]bool)
	// squareSet := make(map[byte]bool)

	// // check each row for duplicate elems
	// for i := 0; i < 9; i++ {
	// 	for j := 0; j < 9; j++ {
	// 		if board[i][j] != '.' {
	// 			if _, exists := rowSet[board[i][j]-'0']; exists {
	// 				return false;
	// 			} else {
	// 				rowSet[board[i][j]-'0'] = true
	// 			}
	// 		}

	// 		if board[j][i] != '.' {
	// 			if _, exists := colSet[board[j][i]-'0']; exists {
	// 				return false
	// 			} else {
	// 				colSet[board[j][i]-'0'] = true
	// 			}
	// 		}
	// 		// unfinished code...
	// 		if board[][] != '.' {
	// 			if _, exists := squareSet[board[][]-'0']; exists {
	// 				return false
	// 			} else {
	// 				squareSet[board[][]-'0'] = true
	// 			}
	// 		}
	// 	}
	// 	fmt.Println(colSet)
	// 	rowSet = make(map[byte]bool)
	// 	colSet = make(map[byte]bool)
	// 	squareSet := make(map[byte]bool)
	// }
	// // check every 3x3 grid for each elem

	// return true

	// ITERATION 2:
	rows := [10][10]bool{} // static 2D (9x9) arrays
	cols := [10][10]bool{}
	squares := [10][10]bool{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := board[i][j] - '0'
			squareIndex := (i/3)*3 + j/3

			if rows[i][num] || cols[j][num] || squares[squareIndex][num] {
				return false
			}

			rows[i][num] = true
			cols[j][num] = true
			squares[squareIndex][num] = true
		}
	}

	return true
}

// @lc code=end

