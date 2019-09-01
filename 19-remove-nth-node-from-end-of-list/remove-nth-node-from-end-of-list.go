package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
    if head == nil || n <= 0 {
        return head
    }
    
    result := &ListNode{
        Next: head,
    }

    l := head
    count := 1
    del := result
    
    for ; l.Next != nil ; {
        
        if count >= n {
            del = del.Next
        }
        
        l = l.Next
        count++
    }
    
    del.Next = del.Next.Next
    
    return result.Next
}