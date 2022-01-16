// Runtime: 8 ms, faster than 88.38% of Go online submissions for Remove Linked List Elements.
// Memory Usage: 4.7 MB, less than 100.00% of Go online submissions for Remove Linked List Elements.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    
    prev := &ListNode{Next: head}
    
    for prev.Next != nil {
        if head.Val == val {
            head = head.Next
        }
        
        if prev.Next.Val == val {
            if prev.Next.Next == nil {
                prev.Next = nil
                return head
            }
            prev.Next = prev.Next.Next
        } else {
            prev = prev.Next
        }
    }
    
    return head
}