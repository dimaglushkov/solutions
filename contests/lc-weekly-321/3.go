package main

import . "github.com/dimaglushkov/solutions/ADS/list"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNodes_(head *ListNode) *ListNode {
	p := head.Next
	mv := head.Val
	cnt := 1
	var res, prev *ListNode

	for p != nil {
		prev = p
		if p.Val > mv {
			cnt = 0
			mv = p.Val
		}
		if p.Val == mv {
			cnt++
		}
		p = p.Next
	}

	for i := 0; i < cnt; i++ {
		if p == nil {
			res = &ListNode{Val: mv, Next: nil}
			p = res
		} else {
			p.Next = &ListNode{Val: mv, Next: nil}
			p = p.Next
		}
	}
	if prev != nil && prev.Val != mv {
		p.Next = prev
	}
	return res
}

func removeNodes(head *ListNode) *ListNode {
	vals := make([]int, 0)
	p := head
	for p != nil {
		vals = append(vals, p.Val)
		p = p.Next
	}

	res := new(ListNode)
	x := vals[len(vals)-1]
	for i := len(vals) - 1; i >= 0; i-- {
		if vals[i] >= x {
			res.Next = &ListNode{Val: vals[i], Next: res.Next}
			x = vals[i]
		}
	}

	return res.Next
}

func main() {
	println(removeNodes(NewList([]int{1})))
	println(removeNodes(NewList([]int{5, 2, 13, 3, 8})))
	println(removeNodes(NewList([]int{1, 1, 1, 1})))
}
