// Runtime: 2 ms, faster than 74.58% of Go online submissions for Intersection of Two Arrays II.
// Memory Usage: 3.5 MB, less than 16.14% of Go online submissions for Intersection of Two Arrays II.

func intersect(nums1 []int, nums2 []int) []int {
    nums1Map := make(map[int]int, 0)
    
    // Solution can be optimized by comparing the length of arrays and adding elements to map from the smaller array. 
    
    for _, num := range nums1 {
        _, ok := nums1Map[num]
        if ok {
            nums1Map[num] += 1
        } else {
            nums1Map[num] = 1
        }
    }
    res := make([]int, 0)
    
    for _, n := range nums2 {
        occurrence, ok := nums1Map[n]
        if ok && occurrence > 0 {
            res = append(res, n)
            nums1Map[n] = occurrence - 1
        }
    }
    
    return res
}