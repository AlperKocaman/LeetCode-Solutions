// Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Linked List.
// Memory Usage: 3.2 MB, less than 9.09% of Go online submissions for Reverse Linked List.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    lstHead, _ := reverseListHelper(nil, head)
    return lstHead
}

func reverseListHelper(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
    if tail == nil {
        return head, head
    }
    
    lstHead, last := reverseListHelper(tail, tail.Next)
    last.Next = head
    return lstHead, head
}