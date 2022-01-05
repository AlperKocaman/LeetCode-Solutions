func containsDuplicate(nums []int) bool {
    
    seenLst := make(map[int]struct{})
    for _, val := range nums {
        if _, ok := seenLst[val]; ok {
            return true
        } 
        seenLst[val] = struct{}{}
    }
    return false
}