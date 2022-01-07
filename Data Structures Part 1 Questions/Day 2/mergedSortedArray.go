// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Sorted Array.
// Memory Usage: 2.5 MB, less than 32.23% of Go online submissions for Merge Sorted Array.

func merge(nums1 []int, m int, nums2 []int, n int)  {
    i := m+n-1
    nums1Idx := m-1
    nums2Idx := n-1
    for nums1Idx >= 0 && nums2Idx >= 0 {
        if nums1[nums1Idx] > nums2[nums2Idx] {
            nums1[i] = nums1[nums1Idx]
            nums1Idx--
        } else {
            nums1[i] = nums2[nums2Idx]
            nums2Idx--
        }
        i--
    }
    
    for nums2Idx >= 0 {
        nums1[nums2Idx] = nums2[nums2Idx]
        nums2Idx--
    } 
}