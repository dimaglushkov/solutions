package main

import (
	"fmt"
)

// source: https://leetcode.com/problems/merge-two-sorted-lists/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(num []int) *ListNode {
	var startNode, prevNode, iNode *ListNode
	if len(num) == 0 {
		return &ListNode{}
	}
	startNode = &ListNode{Val: num[0], Next: nil}
	prevNode = startNode
	for _, i := range num[1:] {
		iNode = &ListNode{Val: i, Next: nil}
		prevNode.Next = iNode
		prevNode = iNode
	}
	return startNode
}

func main() {
	// Example 1
	var list11 *ListNode = buildList([]int{1, 2, 4})
	var list21 *ListNode = buildList([]int{1, 3, 4})
	var res1 = mergeTwoLists(list11, list21)
	fmt.Println("Expected: [1,1,2,3,4,4]	Output: ", res1)

	// Example 2
	var list12 *ListNode = buildList([]int{})
	var list22 *ListNode = buildList([]int{})
	fmt.Println("Expected: []	Output: ", mergeTwoLists(list12, list22))

	// Example 3
	var list13 *ListNode = buildList([]int{})
	var list23 *ListNode = buildList([]int{0})
	fmt.Println("Expected: [0]	Output: ", mergeTwoLists(list13, list23))

}
