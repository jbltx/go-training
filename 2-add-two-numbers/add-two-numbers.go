package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l3 := &ListNode{}
    result := l3
    carry := 0
    for ;l1 != nil || l2 != nil; {
        x, y := 0, 0
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        }
        sum := x + y + carry
        if sum >= 10 {
            sum -= 10
            carry = 1
        } else {
            carry = 0
        }
        l3.Next = &ListNode{
            Val: sum,
        }
        l3 = l3.Next
    }
    if carry > 0 {
        l3.Next = &ListNode{
            Val: carry,
        }
    }
    return result.Next
}