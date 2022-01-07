func maxSubArray(nums []int) int {
    intSize := 32
    max := -1 << (intSize - 1)
    cur := 0
    for _, i := range nums {
        if cur < 0 {
            cur = 0
        }
        cur += i
        if cur > max {
            max = cur
        }
    }
    return max
}