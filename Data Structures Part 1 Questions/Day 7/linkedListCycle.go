// Runtime: 21 ms, faster than 7.87% of Go online submissions for Linked List Cycle.
// Memory Usage: 4.4 MB, less than 56.64% of Go online submissions for Linked List Cycle.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    
    maxNodeVal := 100000 // 10^5
    if head == nil {
        return false
    }
    
    for head.Next != nil {
        if head.Val > maxNodeVal {
            return true
        }
        head.Val = maxNodeVal + 1
        head = head.Next
    }
    return false
}