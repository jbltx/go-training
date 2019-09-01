package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    list := &ListNode{
        Next: head,
    }
    
    for l := list; l.Next != nil && l.Next.Next != nil;   {
        isDup := false
        for ; l.Next.Next != nil && l.Next.Val == l.Next.Next.Val ; {
            isDup = true
            l.Next.Next = l.Next.Next.Next
        }
        if isDup {
            l.Next = l.Next.Next
        } else {
            l=l.Next
        }
    }
    
    return list.Next
}