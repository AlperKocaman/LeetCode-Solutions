// Runtime: 4 ms, faster than 67.36% of Go online submissions for Valid Sudoku.
// Memory Usage: 3.6 MB, less than 30.19% of Go online submissions for Valid Sudoku.

func isValidSudoku(board [][]byte) bool {
    // first 9 for rows, second 9 for cols and last 9 for 3*3 sub-boxes.
    sets := make([]map[byte]struct{}, 27)
    for k:=0 ; k<27 ; k++ {
        sets[k] = make(map[byte]struct{})
    }
    
    for i:=0 ; i<9 ; i++ {
        for j:=0 ; j<9 ; j++ {
            if board[i][j] == '.' {
                continue
            }
            _, ok := sets[i][board[i][j]]
            if ok {
                return false
            }
            sets[i][board[i][j]] = struct{}{}
            _, ok = sets[j+9][board[i][j]]
            if ok {
                return false
            }
             sets[j+9][board[i][j]] = struct{}{}
            _, ok = sets[18+i/3*3+j/3][board[i][j]]
            if ok {
                return false
            }
             sets[18+i/3*3+j/3][board[i][j]] = struct{}{}
        }
    }
    
    return true
}