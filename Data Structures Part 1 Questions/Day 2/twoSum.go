// Runtime: 4 ms, faster than 97.12% of Go online submissions for Two Sum.
// Memory Usage: 4.3 MB, less than 31.34% of Go online submissions for Two Sum.

func twoSum(nums []int, target int) []int {
    deficitMap := make(map[int]int, 0)
    for curIndex, num := range nums {
        if firstIndex, ok := deficitMap[num]; ok {
            return []int{firstIndex, curIndex}
        }
        deficitMap[target-num] = curIndex
        
    }
    
    // this line should not be executed
    return []int{}
}