// Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
// Memory Usage: 2.1 MB, less than 98.97% of Go online submissions for Pascal's Triangle.

func generate(numRows int) [][]int {
    pascalTri := make([][]int, numRows)
    
    for i:=0 ; i<numRows ; i++ {
        pascalTri[i] = make([]int, i+1)
        pascalTri[i][0] = 1
        pascalTri[i][i] = 1
        for j:=1 ; j<i ; j++ {
            pascalTri[i][j] = pascalTri[i-1][j-1] + pascalTri[i-1][j]
        } 
    }
    
    return pascalTri
}