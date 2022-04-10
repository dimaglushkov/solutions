package list

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(num []int) *ListNode {
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

func NewReverseList(num []int) *ListNode {
	var prevNode, iNode *ListNode
	for _, i := range num {
		iNode = &ListNode{Val: i, Next: prevNode}
		prevNode = iNode
	}
	return iNode
}

func (n ListNode) String() string {
	values := make([]string, 0)
	p := &n

	for p != nil {
		values = append(values, strconv.FormatInt(int64(p.Val), 10))
		p = p.Next
	}
	return "[" + strings.Join(values, " ") + "]"
}
