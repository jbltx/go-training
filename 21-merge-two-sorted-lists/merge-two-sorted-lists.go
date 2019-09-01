package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    
    result := &ListNode{}
    curr := result
    
    for ;l1 != nil && l2 != nil; {
        
        if l1.Val < l2.Val {
            curr.Next = &ListNode{
                Val: l1.Val,
            }
            l1 = l1.Next
        } else {
            curr.Next = &ListNode{
                Val: l2.Val,
            }
            l2 = l2.Next
        }
        
        curr = curr.Next
    }
    
    if l1 != nil {
        curr.Next = l1
    } else if l2 != nil {
        curr.Next = l2
    }
    
    return result.Next
}