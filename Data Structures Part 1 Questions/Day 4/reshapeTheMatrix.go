// Runtime: 8 ms, faster than 95.75% of Go online submissions for Reshape the Matrix.
// Memory Usage: 6.4 MB, less than 94.81% of Go online submissions for Reshape the Matrix.

func matrixReshape(mat [][]int, r int, c int) [][]int {
 
    if len(mat) * len(mat[0]) != r*c {
        return mat
    }
    
    res := make([][]int, r)
    for i := range res {
        res[i] = make([]int, c)
    }
    
    r1, c1 := 0, 0
    
    for rowIdx := 0; rowIdx < r ; rowIdx++ {
        for colIdx := 0 ; colIdx < c ; colIdx++ {
            res[rowIdx][colIdx] = mat[r1][c1]
            c1++
            if c1 == len(mat[0]) {
                c1 = 0
                r1++
            }
        }
    }
    return res
}