package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    
    if head == nil {
        return nil
    }
    
    h := head
    
    for ; ; h = h.Next {
                
        for ; h.Next != nil && h.Val == h.Next.Val ; {
            h.Next = h.Next.Next
        }
        
        if h.Next == nil {
            break
        }
    }
    
    return head
}