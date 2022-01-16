// Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
// Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Merge Two Sorted Lists.

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    
    if list2 == nil {
        return list1
    }
    
    var head *ListNode
    var last *ListNode
    
    if list1.Val <= list2.Val {
        head = list1
        last = list1
        if last.Next == nil {
            last.Next = list2
            return head
        }
        list1 = list1.Next
    } else {
        head = list2
        last = list2
        if last.Next == nil {
            last.Next = list1
            return head
        }
        list2 = list2.Next 
    }
    
    for list1 != nil && list2 != nil {
        if list1.Val <= list2.Val {
            last.Next = list1
            list1 = list1.Next
        } else {
            last.Next = list2
            list2 = list2.Next
        }
        last = last.Next
    }
    
    if list1 != nil {
        last.Next= list1
    } else if list2 != nil {
        last.Next= list2
    }
    
    return head
}