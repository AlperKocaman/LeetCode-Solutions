// Runtime: 3 ms, faster than 85.64% of Go online submissions for Remove Duplicates from Sorted List.
// Memory Usage: 3.2 MB, less than 45.36% of Go online submissions for Remove Duplicates from Sorted List.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    
    duplicateChecker := head // front pointer
    prev := &ListNode{Val : 101, Next: head} // back pointer
    
    for duplicateChecker != nil {
        if duplicateChecker.Val != prev.Val {
            if prev.Next != duplicateChecker{
                prev.Next = duplicateChecker
            }
            prev = prev.Next
        }
        duplicateChecker = duplicateChecker.Next
    }
    
    prev.Next = duplicateChecker
    
    return head
    
}