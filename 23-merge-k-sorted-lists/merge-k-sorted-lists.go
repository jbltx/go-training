package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "container/heap"

type Item struct {
	value       *ListNode
	priority    int
	index       int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func mergeKLists(lists []*ListNode) *ListNode {
    
    head := &ListNode{}
    point := head
    pq := make(PriorityQueue, 0)
    heap.Init(&pq)
    
    for _,l := range lists {
        if l != nil {
            heap.Push(&pq, &Item{
                value: l,
                priority: l.Val,
            })
        }
    }
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)
        point.Next = &ListNode{
            Val: item.priority,
        }
        point = point.Next
        item.value = item.value.Next
        if item.value != nil {
            heap.Push(&pq, &Item{
                value: item.value,
                priority: item.value.Val,
            })
        }
    }
    
    return head.Next
}