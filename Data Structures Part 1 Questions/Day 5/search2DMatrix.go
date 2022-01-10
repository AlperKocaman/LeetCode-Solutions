// Runtime: 0 ms, faster than 100.00% of Go online submissions for Search a 2D Matrix.
// Memory Usage: 2.7 MB, less than 37.53% of Go online submissions for Search a 2D Matrix.

func searchMatrix(matrix [][]int, target int) bool {
    rows, cols := len(matrix), len(matrix[0])
    min, mid, max := 0, rows*cols/2, rows*cols-1
    
    for max >= min && mid >= 0 {
        if matrix[mid/cols][mid%cols] == target {
            return true
        } else if matrix[mid/cols][mid%cols] > target {
            max = mid -1
        } else {
            min = mid +1
        }
        mid = (min+max)/2
    }
    return false
}