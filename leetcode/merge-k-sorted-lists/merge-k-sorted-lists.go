package main

import (
	"container/heap"
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
)

// source: https://leetcode.com/problems/merge-k-sorted-lists/

type Heap []*ListNode

func NewHeap(values ...*ListNode) *Heap {
	h := make(Heap, 0, len(values))
	h = append(h, values...)
	return &h
}

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *Heap) Pop() interface{} {
	lastId := len(*h) - 1

	if (*h)[lastId].Next != nil {
		res := (*h)[lastId]
		(*h)[lastId] = (*h)[lastId].Next
		heap.Fix(h, lastId)
		return res
	}
	res := (*h)[lastId]
	*h = (*h)[0:lastId]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	res := new(ListNode)
	cur := res
	h := NewHeap()
	heap.Init(h)

	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)

		}
	}
	for h.Len() > 0 {
		cur.Next = heap.Pop(h).(*ListNode)
		cur = cur.Next
	}

	return res.Next
}

func main() {
	// Example 6
	var lists6 = []*ListNode{NewList([]int{1, 2, 3}), NewList([]int{4, 5, 6, 7})}
	fmt.Println("Expected: [1,2,3,4,5,6,7]	Output: ", mergeKLists(lists6))

	// Example 5
	var lists5 = []*ListNode{NewList([]int{1, 2, 2}), NewList([]int{1, 1, 2})}
	fmt.Println("Expected: [1,1,1,2,2,2]	Output: ", mergeKLists(lists5))

	// Example 4
	var lists4 []*ListNode = []*ListNode{nil}
	fmt.Println("Expected: []	Output: ", mergeKLists(lists4))

	// Example 3
	var lists3 []*ListNode = []*ListNode{{}}
	fmt.Println("Expected: [0]	Output: ", mergeKLists(lists3))

	// Example 1
	var lists1 = []*ListNode{NewList([]int{1, 4, 5}), NewList([]int{1, 3, 4}), NewList([]int{2, 6})}
	fmt.Println("Expected: [1,1,2,3,4,4,5,6]	Output: ", mergeKLists(lists1))

	// Example 2
	var lists2 []*ListNode = nil
	fmt.Println("Expected: []	Output: ", mergeKLists(lists2))
}
